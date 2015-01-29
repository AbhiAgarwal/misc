// http://rosettacode.org/wiki/Dijkstra%27s_algorithm#Go
package graph

import (
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := []Edge{
		{"a", "b", 7},
		{"a", "c", 9},
		{"a", "f", 14},
		{"b", "c", 10},
		{"b", "d", 15},
		{"c", "d", 11},
		{"c", "f", 2},
		{"d", "e", 6},
		{"e", "f", 9},
	}
	directed := true
	start := "a"
	end := "e"
	findAll := false

	// construct linked representation of example data
	allNodes, startNode, endNode := LinkGraph(graph, directed, start, end)
	if len(allNodes) != 6 {
		t.Error("Wrong number of All Nodes")
	}
	if len(graph) != 9 {
		t.Error("Wrong number in Graph")
	}

	if startNode == nil {
		t.Error("start node %q not found in graph\n", start)
	}
	if findAll {
		endNode = nil
	} else if endNode == nil {
		t.Error("end node %q not found in graph\n", end)
	}

	// run Dijkstra's shortest path algorithm
	paths := Dijkstra(allNodes, startNode, endNode)
	answer := []string{"a", "c", "d", "e"}
	for _, p := range paths {
		for x, y := range p.path {
			if y != answer[x] {
				t.Error("Wrong answer!")
			}
		}
		if p.length == 4 {
			t.Error("Length of path is incorrect")
		}
	}
}
