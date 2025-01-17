package orchestrator

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/streamingfast/substreams"
	"github.com/streamingfast/substreams/block"
	pbsubstreams "github.com/streamingfast/substreams/pb/sf/substreams/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type WorkerPool struct {
	workers chan *Worker
}

func NewWorkerPool(workerCount int, originalRequestModules *pbsubstreams.Modules, grpcClientFactory func() (pbsubstreams.StreamClient, []grpc.CallOption, error)) *WorkerPool {
	zlog.Info("initiating worker pool", zap.Int("worker_count", workerCount))
	workers := make(chan *Worker, workerCount)
	for i := 0; i < workerCount; i++ {
		workers <- &Worker{
			originalRequestModules: originalRequestModules,
			grpcClientFactory:      grpcClientFactory,
		}
	}
	return &WorkerPool{
		workers: workers,
	}
}

func (p *WorkerPool) Borrow() *Worker {
	w := <-p.workers
	return w
}

func (p *WorkerPool) ReturnWorker(worker *Worker) {
	p.workers <- worker
}

type Worker struct {
	grpcClientFactory      func() (pbsubstreams.StreamClient, []grpc.CallOption, error)
	originalRequestModules *pbsubstreams.Modules
}

func (w *Worker) Run(ctx context.Context, job *Job, respFunc substreams.ResponseFunc) ([]*block.Range, error) {
	start := time.Now()
	zlog.Info("running job", zap.Object("job", job))
	defer func() {
		zlog.Info("job completed", zap.Object("job", job), zap.Duration("in", time.Since(start)))
	}()

	grpcClient, grpcCallOpts, err := w.grpcClientFactory()
	if err != nil {
		zlog.Error("getting grpc client", zap.Error(err))
		return nil, fmt.Errorf("grpc client factory: %w", err)
	}

	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"substreams-partial-mode": "true"}))

	request := job.createRequest(w.originalRequestModules)

	stream, err := grpcClient.Blocks(ctx, request, grpcCallOpts...)
	if err != nil {
		return nil, fmt.Errorf("getting block stream: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			zlog.Warn("context cancel will waiting for stream data, worker is terminating")
			return nil, ctx.Err()
		default:
		}

		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				zlog.Info("worker done", zap.Object("job", job))
				trailers := stream.Trailer().Get("substreams-partials-written")
				var partialsWritten []*block.Range
				if len(trailers) != 0 {
					zlog.Info("partial written", zap.String("trailer", trailers[0]))
					partialsWritten = block.ParseRanges(trailers[0])
				}

				return partialsWritten, nil
			}
			zlog.Warn("worker done on stream error", zap.Error(err))
			return nil, fmt.Errorf("receiving stream resp: %w", err)
		}

		switch r := resp.Message.(type) {
		case *pbsubstreams.Response_Progress:
			err := respFunc(resp)
			if err != nil {
				zlog.Warn("worker done on respFunc error", zap.Error(err))
				return nil, fmt.Errorf("sending progress: %w", err)
			}
		case *pbsubstreams.Response_SnapshotData:
			_ = r.SnapshotData
		case *pbsubstreams.Response_SnapshotComplete:
			_ = r.SnapshotComplete
		case *pbsubstreams.Response_Data:
			// These are not returned by virtue of `returnOutputs`
		}
	}
}
