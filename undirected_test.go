package graph

import "testing"

func TestUnDirectedGraph_AddVertexAndNumVertices(t *testing.T) {
	g := NewUnDirectedAdjacencyList()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("A")

	if g.NumVertices() != 2 {
		t.Errorf("Expected 2 vertices, got %d", g.NumVertices())
	}
}

func TestUnDirectedGraph_AddUndirectedEdgeAndNeighbors(t *testing.T) {
	g := NewUnDirectedAdjacencyList()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddUndirectedEdge("A", "B", 4.2)

	if g.NumEdges() != 1 {
		t.Errorf("Expected 1 edge, got %d", g.NumEdges())
	}

	neighborsA := g.Neighbors("A")
	if len(neighborsA) != 1 || neighborsA[0] != "B" {
		t.Errorf("Expected B to be neighbor of A")
	}

	neighborsB := g.Neighbors("B")
	if len(neighborsB) != 1 || neighborsB[0] != "A" {
		t.Errorf("Expected A to be neighbor of B")
	}
}

func TestUnDirectedGraph_BFS(t *testing.T) {
	g := NewUnDirectedAdjacencyList()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddUndirectedEdge("A", "B", 1.0)
	g.AddUndirectedEdge("B", "C", 1.0)

	distances := g.BFS("A")
	if distances["A"] != 0 || distances["B"] != 1 || distances["C"] != 2 {
		t.Errorf("BFS distances incorrect: %v", distances)
	}
}

func TestUnDirectedGraph_DFS(t *testing.T) {
	g := NewUnDirectedAdjacencyList()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddUndirectedEdge("A", "B", 1.0)

	visited := g.DFS("A")
	if !visited["A"] || !visited["B"] || visited["C"] {
		t.Errorf("DFS traversal incorrect: %v", visited)
	}
}

func TestUnDirectedGraph_UCC(t *testing.T) {
	g := NewUnDirectedAdjacencyList()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")

	g.AddUndirectedEdge("A", "B", 1.0)
	g.AddUndirectedEdge("C", "D", 1.0)

	components := g.UCC()

	if components["A"] != components["B"] {
		t.Errorf("A and B must be in the same component")
	}
	if components["C"] != components["D"] {
		t.Errorf("C and D must be in the same component")
	}
	if components["A"] == components["C"] || components["A"] == components["E"] || components["C"] == components["E"] {
		t.Errorf("A, C and E must be in three different components")
	}
}
