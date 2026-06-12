package graph

import (
	"slices"
	"testing"
)

func TestDirected1(t *testing.T) {
	// for logging messages:
	//   go test -v -run TestDirected1

	var directedGraph DirectedGraph

	directedGraph = NewDirectedAdjacencyList()

	vertices := []string{"A", "B", "C", "D", "E", "F"}
	for _, v := range vertices {
		directedGraph.AddVertex(v)
	}
	for _, v := range vertices[1:5] {
		directedGraph.AddDirectedEdge(vertices[0], v, 1.)
	}
	directedGraph.AddDirectedEdge(vertices[4], vertices[0], 1.)
	directedGraph.AddDirectedEdge(vertices[5], vertices[0], 1.)
	directedGraph.AddDirectedEdge(vertices[3], vertices[2], 1.)

	// example of logging messages
	t.Logf("Number of Vertices: %v", directedGraph.NumVertices())
	t.Logf("Number of Edges: %v", directedGraph.NumEdges())
	for _, nodeId := range directedGraph.Successors("A") {
		t.Logf("A has successor: %v", nodeId)
	}
	for _, nodeId := range directedGraph.Predecessors("A") {
		t.Logf("A has predecessor: %v", nodeId)
	}

	//
	if directedGraph.NumVertices() != len(vertices) {
		t.Fatalf("NumVertices failed!")
	}
	if directedGraph.NumEdges() != 7 {
		t.Fatalf("NumEdges failed!")
	}

	suc := directedGraph.Successors("A")
	for _, v := range vertices[1:5] {
		if !slices.Contains(suc, v) {
			t.Fatalf("%v not in Successors of A", v)
		}
	}
	pre := directedGraph.Predecessors("A")
	if !(slices.Contains(pre, "E") && slices.Contains(pre, "F")) {
		t.Fatalf("E or F not in Predessors of A")
	}

}
