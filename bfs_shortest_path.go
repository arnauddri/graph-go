package graph

import (
	"fmt"
)

func bfsShortestPath(g *Graph, start VertexId) {
	queue := []VertexId{start}
	visited := make(map[VertexId]bool)
	dist := make(map[VertexId]int)
	var next []VertexId

	for len(queue) > 0 {
		next = []VertexId{}
		for _, vertex := range queue {
			visited[vertex] = true
			neighbours := g.GetNeighbours(vertex).VertexesIter()

			for neighbour := range neighbours {

				ok, _ := visited[neighbour]
				if !ok {
					dist[neighbour] = dist[vertex] + 1
					next = append(next, neighbour)
				}
			}
		}
		queue = next
	}
	fmt.Println(g)
	fmt.Println(dist)
}
