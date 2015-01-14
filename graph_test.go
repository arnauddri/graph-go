package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph()

	for i := 0; i < 10; i++ {
		v := VertexId(i)
		g.AddVertex(v)
	}

	if len(g.edges) != 10 {
		fmt.Println(g)
		t.Error()
	}

	for i := 0; i < 10; i++ {
		g.AddEdge(VertexId(i), VertexId(i%2))
	}

	if g.isEdge(0, 8) == false || g.isEdge(0, 9) == true {
		fmt.Println(g)
		t.Error()
	}

	g.RemoveVertex(VertexId(9))

	if g.isVertex(VertexId(9)) {
		fmt.Println(g.edges[9] == nil)
		t.Error()
	}

	g.RemoveEdge(0, 8)

	if g.isEdge(VertexId(0), VertexId(8)) || g.edgesCount != 7 {
		fmt.Println(g.isEdge(VertexId(0), VertexId(8)), g.edgesCount)
		t.Error()
	}

	c := g.EdgesIter()
	opened := true

	countEdge := 0
	for _ = range c {
		countEdge++
	}

	if g.EdgesCount() != countEdge && !opened {
		fmt.Println(countEdge, g.edges)
		t.Error()
	}

	d := g.VertexesIter()
	countVertices := 0

	for _ = range d {
		countVertices++
	}

	if countVertices != len(g.edges) {
		fmt.Println(countVertices, g.edges)
		t.Error()
	}
}

func TestBfs(t *testing.T) {
	h := NewGraph()

	for i := 0; i < 10; i++ {
		v := VertexId(i)
		h.AddVertex(v)
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(VertexId(i), VertexId(i+1))
	}

	bfsShortestPath(h, VertexId(1))
	bfs(h, VertexId(2))
}
