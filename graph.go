package graph

import (
	"sync"
)

type Vertex struct {
	key       string
	neighbors []*Vertex
	sync.RWMutex
}

type Graph struct {
	vertexes []Vertex
	sync.RWMutex
}

func (v *Vertex) GetNeighbors() []*Vertex {
	if v == nil {
		return nil
	}

	v.RLock()
	neighbors := v.neighbors
	v.RUnlock()

	return neighbors
}

func (v *Vertex) Key() string {
	if v == nil {
		return ""
	}

	v.RLock()
	key := v.key
	v.RUnlock()

	return key
}

func (g *Graph) Add(v *Vertex) bool {
	g.Lock()
	defer g.Unlock()

	g.vertexes = append(g.vertexes, *v)

	return true
}
