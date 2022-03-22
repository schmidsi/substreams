package manifest

import (
	"fmt"
	"sort"

	pbtransform "github.com/streamingfast/substreams/pb/sf/substreams/transform/v1"
	"github.com/yourbasic/graph"
)

type ModuleGraph struct {
	*graph.Mutable

	modules     []*pbtransform.Module
	moduleIndex map[string]int
	indexIndex  map[int]*pbtransform.Module
}

func NewModuleGraph(modules []*pbtransform.Module) (*ModuleGraph, error) {
	g := &ModuleGraph{
		Mutable:     graph.New(len(modules)),
		modules:     modules,
		moduleIndex: make(map[string]int),
		indexIndex:  make(map[int]*pbtransform.Module),
	}

	for i, module := range modules {
		g.moduleIndex[module.Name] = i
		g.indexIndex[i] = module
	}

	for i, module := range modules {
		for _, input := range module.Inputs {

			var moduleName string
			if v := input.GetMap(); v != nil {
				moduleName = v.ModuleName
			} else if v := input.GetStore(); v != nil {
				moduleName = v.ModuleName
			}
			if moduleName == "" {
				continue
			}

			if j, found := g.moduleIndex[moduleName]; found {
				g.AddCost(i, j, 1)
			}
		}
	}

	if !graph.Acyclic(g) {
		return nil, fmt.Errorf("modules graph has a cycle")
	}

	return g, nil
}

func (g *ModuleGraph) topSort() ([]*pbtransform.Module, bool) {
	order, ok := graph.TopSort(g)
	if !ok {
		return nil, ok
	}

	var res []*pbtransform.Module
	for _, i := range order {
		res = append(res, g.indexIndex[i])
	}

	return res, ok
}

func (g *ModuleGraph) AncestorsOf(moduleName string) ([]*pbtransform.Module, error) {
	if _, found := g.moduleIndex[moduleName]; !found {
		return nil, fmt.Errorf("could not find module %s in graph", moduleName)
	}

	_, distances := graph.ShortestPaths(g, g.moduleIndex[moduleName])

	var res []*pbtransform.Module
	for i, d := range distances {
		if d >= 1 {
			res = append(res, g.indexIndex[i])
		}
	}

	return res, nil
}

func (g *ModuleGraph) AncestorStoresOf(moduleName string) ([]*pbtransform.Module, error) {
	ancestors, err := g.AncestorsOf(moduleName)
	if err != nil {
		return nil, err
	}

	result := make([]*pbtransform.Module, 0, len(ancestors))
	for _, a := range ancestors {
		if a.GetKindStore() != nil {
			result = append(result, a)
		}
	}

	return result, nil
}

func (g *ModuleGraph) ParentsOf(moduleName string) ([]*pbtransform.Module, error) {
	if _, found := g.moduleIndex[moduleName]; !found {
		return nil, fmt.Errorf("could not find module %s in graph", moduleName)
	}

	_, distances := graph.ShortestPaths(g, g.moduleIndex[moduleName])

	var res []*pbtransform.Module
	for i, d := range distances {
		if d == 1 {
			res = append(res, g.indexIndex[i])
		}
	}

	return res, nil
}

func (g *ModuleGraph) StoresDownTo(moduleName string) ([]*pbtransform.Module, error) {
	if _, found := g.moduleIndex[moduleName]; !found {
		return nil, fmt.Errorf("could not find module %s in graph", moduleName)
	}

	_, distances := graph.ShortestPaths(g, g.moduleIndex[moduleName])

	var res []*pbtransform.Module
	for i, d := range distances {
		if d >= 0 { // connected node or myself
			module := g.indexIndex[i]
			if module.GetKindStore() != nil {
				res = append(res, g.indexIndex[i])
			}
		}
	}

	return res, nil
}

func (g *ModuleGraph) ModulesDownTo(moduleName string) ([]*pbtransform.Module, error) {
	if _, found := g.moduleIndex[moduleName]; !found {
		return nil, fmt.Errorf("could not find module %s in graph", moduleName)
	}

	_, distances := graph.ShortestPaths(g, g.moduleIndex[moduleName])

	var res []*pbtransform.Module
	for i, d := range distances {
		if d >= 0 { // connected node or myself
			res = append(res, g.indexIndex[i])
		}
	}

	return res, nil
}

func (g *ModuleGraph) GroupedModulesDownTo(moduleName string) ([][]*pbtransform.Module, error) {
	v, found := g.moduleIndex[moduleName]

	if !found {
		return nil, fmt.Errorf("could not find module %s in graph", moduleName)
	}

	mods, err := g.ModulesDownTo(moduleName)
	if err != nil {
		return nil, fmt.Errorf("could not determine dependencies graph for %s: %w", moduleName, err)
	}

	_, dist := graph.ShortestPaths(g, v)

	distmap := map[int][]*pbtransform.Module{}
	distkeys := []int{}
	for _, mod := range mods {
		mix := g.moduleIndex[mod.Name]
		if _, found := distmap[int(dist[mix])]; !found {
			distkeys = append(distkeys, int(dist[mix]))
		}
		distmap[int(dist[mix])] = append(distmap[int(dist[mix])], mod)
	}

	//reverse sort
	sort.Slice(distkeys, func(i, j int) bool {
		return distkeys[j] < distkeys[i]
	})

	res := make([][]*pbtransform.Module, 0, len(distmap))
	for _, ix := range distkeys {
		res = append(res, distmap[ix])
	}

	return res, nil
}
