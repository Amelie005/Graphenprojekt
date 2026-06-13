package graph

import "math"

//definieren, was eine Kante ist: Ziel und Länge
type Edge struct {
	To     string
	Length float64
}

type MyDirectedGraph struct {
	vertices map[string]bool
	//für jeden Knoten speichern von einer Liste seiner ausgehenden Kanten
	adjList map[string][]Edge
	//für jeden Knoten speichern wir direkt seine Vorgänger-Knoten (also invertiert)
	invList map[string][]string
}

//Quasi eine Konstruktor Funktion

func NewDirectedAdjacencyList() DirectedGraph {
	return &MyDirectedGraph{ //&: nicht gesamter Praph zurückgegeben, sondern Pointer auf den Graphen
		vertices: make(map[string]bool),     //aktiviert die Knotenliste
		adjList:  make(map[string][]Edge),   //aktiviert die Kantenliste (Nachfolger)
		invList:  make(map[string][]string), //initialisieren der invertierten Liste (Vorgänger)
	}
}

//fügt einen Knoten zum Graphen hinzu
func (g *MyDirectedGraph) AddVertex(nodeId string) {
	//wenn der Knoten noch nicht existiert, wird er angelegt
	if _, exists := g.vertices[nodeId]; !exists {
		g.vertices[nodeId] = true
		g.adjList[nodeId] = []Edge{}   //leere Kantenliste für diesen Knoten erstellen
		g.invList[nodeId] = []string{} //leere Vorgängerliste erstellen
	}
}

//Gibt die Anzahl der Knoten in diesem Graphen zurück
func (g *MyDirectedGraph) NumVertices() int {
	return len(g.vertices) //Länge der vertices Map zurück gegeben
}

//gerichtete Kanten hinzufügen
func (g *MyDirectedGraph) AddDirectedEdge(nodeId1, nodeId2 string, length float64) { //von node 1 zu node 2 mit Länge
	//holen von der aktuellen Kantenliste von Knoten 1 und neue Kante zufügen
	g.adjList[nodeId1] = append(g.adjList[nodeId1], Edge{To: nodeId2, Length: length})

	//Knoten 1 als Vorgänger von Knoten 2 registrieren
	g.invList[nodeId2] = append(g.invList[nodeId2], nodeId1)
}

//Gibt die Anzahl der Kanten in dem Graphen zurück
func (g *MyDirectedGraph) NumEdges() int {
	count := 0
	//durch alle Kantenlisten im Graphen laufen
	for _, edges := range g.adjList {
		count += len(edges) //Anzahl der Kanten addieren
	}
	return count
}

//Gibt die Nachfolger eines Knotens aus
func (g *MyDirectedGraph) Successors(nodeId string) []string {
	edges := g.adjList[nodeId]
	list := make([]string, len(edges))

	//aus jeder Kante das Ziel (also to) heraus holen
	for i, edge := range edges {
		list[i] = edge.To
	}
	return list
}

//Gibt die Vorgänger eines Knotens aus
func (g *MyDirectedGraph) Predecessors(nodeId string) []string {
	//man muss den ganzen Graphen nicht immer wieder nach passenden Knoten durchsuchen,
	//denn die Vorgänger wurden ja schon in einer Liste gespeichert
	return g.invList[nodeId]
}

//Breadth-First Search
//berechnet Abstände (Distanzen) vom Startknoten zu allen anderen Knoten im Graphen
func (g *MyDirectedGraph) BFS(nodeId string) map[string]int {
	//Distanzen in Map gespeichert und ob der Knoten schon besucht wurde
	distances := make(map[string]int)

	//falls der Startknoten gar nicht im Graphen existiert, abbrechen
	if !g.vertices[nodeId] {
		return distances
	}

	//Queue erstellen
	q := NewQueue()

	//Startknoten initialisieren
	distances[nodeId] = 0
	q.Enqueue(nodeId)

	//solange noch Knoten in der Warteschlange sind
	for !q.IsEmpty() {
		//immer den nächsten Knoten herausholen
		current, _ := q.Dequeue()
		currentDist := distances[current]

		//alle Nachfolger des aktuellen Knotens untersuchen
		for _, successor := range g.Successors(current) {
			//wenn der Nachfolger noch keine Distanz hat, wurde er noch nicht besucht
			if _, visited := distances[successor]; !visited {
				//Distanz ist die Entfernung des aktuellen Knotens + 1
				distances[successor] = currentDist + 1
				//Nachfolger in die Queue packen, um später dessen Nachbarn zu prüfen
				q.Enqueue(successor)
			}
		}
	}

	return distances
}

//Depth-First Search
func (g *MyDirectedGraph) DFS(nodeId string) map[string]bool {
	//hier wird gemerkt, welche Knoten schon erreicht wurden
	visited := make(map[string]bool)

	//falls der Startknoten nicht existiert, direkt leere Map zurückgeben
	if !g.vertices[nodeId] {
		return visited
	}

	//rekursive Suche beim Startknoten starten
	g.dfsHelper(nodeId, visited)

	return visited
}

//Hilfsfunktion
//macht die Rekursion bei DFS
func (g *MyDirectedGraph) dfsHelper(current string, visited map[string]bool) {
	//aktuellen Knoten als besucht markieren
	visited[current] = true

	//alle Nachfolger ansehen
	for _, successor := range g.Successors(current) {
		//wenn Nachfolger noch nicht besucht wurde
		if !visited[successor] {
			//dann hier rekursiv ansetzen usw.
			g.dfsHelper(successor, visited)
		}
	}
}

//Shortest-Path Algorithm
//findet den kürzesten Pfad zu einem Knoten in einem gewichteten Graphen
func (g *MyDirectedGraph) Dijkstra(id string) map[string]float64 {
	distances := make(map[string]float64)

	//alle Knoten im Graphen auf "unendlich weit weg" setzen
	for v := range g.vertices {
		distances[v] = math.MaxFloat64
	}

	//falls der Startknoten gar nicht existiert, direkt abbrechen
	if !g.vertices[id] {
		return distances
	}

	//Startknoten initialisieren
	distances[id] = 0.0
	heap := NewMinHeap() //minHeap mit kleinstem Element eben ganz oben
	heap.Push(HeapItem{Id: id, Distance: 0.0})

	//Hauptschleife des Algorithmus
	for !heap.IsEmpty() { //wenn Heap nicht leer ist
		current, _ := heap.Pop() //kleinstes Element poppen

		//wenn wir einen veralteten (längeren) Eintrag aus dem Heap ziehen,
		//ignorieren wir ihn einfach
		//spart Laufzeit
		if current.Distance > distances[current.Id] {
			continue
		}

		//durch alle ausgehenden Kanten des aktuellen Knotens gehen
		for _, edge := range g.adjList[current.Id] {
			//Distanz für neuen potenziellen Weg errechnen
			newDist := distances[current.Id] + edge.Length

			//wenn der neue Weg kürzer ist als der alte bekannte Weg, den neuen Weg wählen
			if newDist < distances[edge.To] {
				distances[edge.To] = newDist
				//Nachfolgeknoten mit der neuen Distanz in den Heap pushen
				heap.Push(HeapItem{Id: edge.To, Distance: newDist})
			}
		}
	}

	//Map säubern
	//Knoten, die unendlich weit weg geblieben sind (bspw.weil sie nicht erreichbar waren),
	//fliegen aus dem Ergebnis raus.
	for v, dist := range distances {
		if dist == math.MaxFloat64 {
			delete(distances, v)
		}
	}

	return distances //Distanzen wiedergeben
}
