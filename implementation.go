package graph

//Definieren, was eine Kante ist: Ziel (To) und Länge (Length)
type Edge struct {
	To     string
	Length float64
}

type MyDirectedGraph struct {
	vertices map[string]bool
	//für jeden Knoten speichern von einer Liste seiner ausgehenden Kanten
	adjList map[string][]Edge
}

func NewDirectedAdjacencyList() DirectedGraph {
	return &MyDirectedGraph{
		vertices: make(map[string]bool),
		adjList:  make(map[string][]Edge),
	}
}

func (g *MyDirectedGraph) AddVertex(nodeId string) {
	//wenn der Knoten noch nicht existiert, wird er angelegt
	if _, exists := g.vertices[nodeId]; !exists {
		g.vertices[nodeId] = true
		g.adjList[nodeId] = []Edge{} //leere Kantenliste für diesen Knoten erstellen
	}
}

func (g *MyDirectedGraph) NumVertices() int {
	return len(g.vertices)
}

//Kanten hinzufügen
func (g *MyDirectedGraph) AddDirectedEdge(nodeId1, nodeId2 string, length float64) {
	//holen von der aktuellen Kantenliste von Knoten 1 und neue Kante zufügen
	g.adjList[nodeId1] = append(g.adjList[nodeId1], Edge{To: nodeId2, Length: length})
}

func (g *MyDirectedGraph) NumEdges() int {
	count := 0
	//durch alle Kantenlisten im Graphen laufen
	for _, edges := range g.adjList {
		count += len(edges) //Anzahl der Kanten addieren
	}
	return count
}

func (g *MyDirectedGraph) Successors(nodeId string) []string {
	edges := g.adjList[nodeId]
	list := make([]string, len(edges))

	//aus jeder Kante das Ziel ("To") heraus holen
	for i, edge := range edges {
		list[i] = edge.To
	}
	return list
}

// Ab hier: Das fehlt uns noch für den Test
func (g *MyDirectedGraph) Predecessors(nodeId string) []string { panic("Predecessors fehlt noch") }

// Die restlichen Methoden, die der Test aktuell noch gar nicht aufruft
func (g *MyDirectedGraph) BFS(nodeId string) map[string]int      { panic("noch nicht eingebaut") }
func (g *MyDirectedGraph) DFS(nodeId string) map[string]bool     { panic("noch nicht eingebaut") }
func (g *MyDirectedGraph) Dijkstra(id string) map[string]float64 { panic("noch nicht eingebaut") }
