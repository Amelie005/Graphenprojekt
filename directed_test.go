package graph

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"
)

//vorgegebene Tests

func TestDirected1(t *testing.T) {
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

	t.Logf("Number of Vertices: %v", directedGraph.NumVertices())
	t.Logf("Number of Edges: %v", directedGraph.NumEdges())
	for _, nodeId := range directedGraph.Successors("A") {
		t.Logf("A has successor: %v", nodeId)
	}
	for _, nodeId := range directedGraph.Predecessors("A") {
		t.Logf("A has predecessor: %v", nodeId)
	}

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

func TestBFSLayers(t *testing.T) {
	g := NewDirectedAdjacencyList()

	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")

	g.AddDirectedEdge("A", "B", 1.0)
	g.AddDirectedEdge("B", "C", 1.0)

	res := g.BFS("A")

	if res["A"] != 0 {
		t.Fatalf("Distanz zu A sollte 0 sein, ist aber %v", res["A"])
	}
	if res["B"] != 1 {
		t.Fatalf("Distanz zu B sollte 1 sein, ist aber %v", res["B"])
	}
	if res["C"] != 2 {
		t.Fatalf("Distanz zu C sollte 2 sein, ist aber %v", res["C"])
	}

	if _, exists := res["D"]; exists {
		t.Fatalf("Knoten D sollte nicht erreichbar sein!")
	}
}

func TestDFSErachability(t *testing.T) {
	g := NewDirectedAdjacencyList()

	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("X")

	g.AddDirectedEdge("A", "B", 1.0)
	g.AddDirectedEdge("B", "C", 1.0)
	g.AddDirectedEdge("A", "D", 1.0)
	g.AddDirectedEdge("D", "A", 1.0)

	res := g.DFS("A")

	if !res["A"] || !res["B"] || !res["C"] || !res["D"] {
		t.Fatalf("A, B, C und D sollten von A aus erreichbar sein!")
	}

	if res["X"] {
		t.Fatalf("Knoten X sollte von A aus absolut nicht erreichbar sein!")
	}
}

func TestDijkstraShortestPath(t *testing.T) {
	g := NewDirectedAdjacencyList()

	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("X")

	g.AddDirectedEdge("A", "B", 1.5)
	g.AddDirectedEdge("B", "C", 2.0)
	g.AddDirectedEdge("A", "C", 10.0)

	res := g.Dijkstra("A")

	if res["C"] != 3.5 {
		t.Fatalf("Dijkstra hat nicht den kürzesten Weg gefunden! Erwartet: 3.5, Got: %v", res["C"])
	}
	if res["B"] != 1.5 {
		t.Fatalf("Erwartet Distanz zu B: 1.5, Got: %v", res["B"])
	}
	if _, exists := res["X"]; exists {
		t.Fatalf("Knoten X sollte nicht erreichbar sein!")
	}
}

//Hilfsfunktionen

func initWebgraph(t *testing.T, webgraph DirectedGraph) {
	file, err := os.Open("./testdata/web-Google.txt")
	if err != nil {
		t.Skip("web-Google.txt nicht gefunden im Ordner ./testdata/! Überspringe Test.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	start := time.Now()
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		if len(fields) > 0 && !strings.HasPrefix(fields[0], "#") && len(fields) == 2 {
			webgraph.AddVertex(fields[0])
			webgraph.AddVertex(fields[1])
			webgraph.AddDirectedEdge(fields[0], fields[1], 1.)
		}
		if (i % 1000000) == 0 {
			elapsed := time.Since(start)
			t.Logf("last took %s\n", elapsed)
			t.Logf("progress: %v\n", i)
			start = time.Now()
		}
		i++
	}
	t.Logf("%v lines processed\n", i)
}

func initGraph9(filename string, graph DirectedGraph) {
	file, err := os.Open(filename)
	if err != nil {
		errorString := fmt.Sprintf("Can't open file %s", filename)
		panic(errorString)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		if len(fields) == 0 {
			continue
		}
		id1 := fields[0]
		graph.AddVertex(id1)
		for _, x := range fields[1:] {
			f := strings.Split(x, ",")
			var length float64
			if l, err := strconv.ParseFloat(f[1], 64); err == nil {
				length = l
			} else {
				panic("convert str2float failed!")
			}
			graph.AddVertex(f[0])

			//da problem9 ungerichtet ist, beide Richtungen eingefügt
			graph.AddDirectedEdge(id1, f[0], length)
			graph.AddDirectedEdge(f[0], id1, length)
		}
	}
}

//meine Tests

// testet den Google Webgraphen (wenn die Datei existiert)
func TestWebGoogleGraph(t *testing.T) {
	graph := NewDirectedAdjacencyList()
	initWebgraph(t, graph)

	if graph.NumVertices() > 0 {
		t.Logf("Google-Graph erfolgreich getestet! Knoten: %v, Kanten: %v", graph.NumVertices(), graph.NumEdges())
	}
}

// prüft, ob Dijkstra auf Problem 9.8 Test das korrekte Ergebnis liefert
func TestProblem9_8_SanityCheck(t *testing.T) {
	graph := NewDirectedAdjacencyList()

	initGraph9("./testdata/problem9.8test.txt", graph)

	distances := graph.Dijkstra("1")

	//erwartete Entfernungen von Knoten 1 zu Knoten 1 bis 8
	expected := []float64{0, 1, 2, 3, 4, 4, 3, 2}

	for i, exp := range expected {
		nodeId := strconv.Itoa(i + 1)
		if distances[nodeId] != exp {
			t.Fatalf("Check fehlgeschlagen bei Knoten %s: Erwartet %v, aber Got %v", nodeId, exp, distances[nodeId])
		}
	}
	t.Log("Check für Problem 9.8 erfolgreich!")
}

// berechnet Ergebnisse für Problem 9.8
func TestProblem9_8_Challenge(t *testing.T) {
	graph := NewDirectedAdjacencyList()

	initGraph9("./testdata/problem9.8.txt", graph)

	distances := graph.Dijkstra("1")

	//die 10 gesuchten Zielknoten
	targets := []string{"7", "37", "59", "82", "99", "115", "133", "165", "188", "197"}

	t.Log("Ergebnisse Dijkstra:")
	for _, target := range targets {
		t.Logf("Kürzeste Distanz von 1 zu %s: %v", target, distances[target])
	}
}
