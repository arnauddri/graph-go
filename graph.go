package graph

import (
	"errors"
)

type VertexId uint

type Vertexes []VertexId

type Edge struct {
	Tail VertexId
	Head VertexId
}

type EdgesIterable interface {
	EdgesIter() <-chan Edge
}

type VertexesIterable interface {
	VertexesIter() <-chan Edge
}

type Graph struct {
	edges      map[VertexId]map[VertexId]bool
	edgesCount int
}

func (g *Graph) EdgesIter() <-chan Edge {
	ch := make(chan Edge)
	go func() {
		for from, connectedVertexes := range g.edges {
			for to, _ := range connectedVertexes {
				if from < to {
					ch <- Edge{from, to}
				}
			}
		}
		close(ch)
	}()
	return ch
}

func (g *Graph) VertexesIter() <-chan VertexId {
	ch := make(chan VertexId)
	go func() {
		for vertex, _ := range g.edges {
			ch <- vertex
		}
		close(ch)
	}()
	return ch
}

func (g *Graph) AddVertex(vertex VertexId) error {
	i, _ := g.edges[vertex]
	if i != nil {
		return errors.New("Vertex already exists")
	}

	g.edges[vertex] = make(map[VertexId]bool)

	return nil
}

func (g *Graph) RemoveVertex(vertex VertexId) error {
	i, _ := g.edges[vertex]
	if i == nil {
		return errors.New("Unknown vertex")
	}

	g.edges[vertex] = nil
	for _, connectedVertexes := range g.edges {
		connectedVertexes[vertex] = false
	}

	return nil
}

func (g *Graph) AddEdge(from, to VertexId) error {
	i, _ := g.edges[from][to]
	j, _ := g.edges[to][from]

	if i == true || j == true {
		return errors.New("Edge already defined")
	}

	g.edges[from][to] = true
	g.edges[to][from] = true

	g.edgesCount++

	return nil
}

func (g *Graph) RemoveEdge(from, to VertexId) error {
	i, _ := g.edges[from][to]
	j, _ := g.edges[to][from]

	if i == false || j == false {
		return errors.New("Edge doesn't exist")
	}

	g.edges[from][to] = false
	g.edges[to][from] = false

	g.edgesCount--

	return nil
}

func (g *Graph) EdgesCount() int {
	return g.edgesCount
}
