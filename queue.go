package graph

//Queue für Strings (da die Knoten-IDs Strings sind)
type Queue struct {
	elements []string
}

//NewQueue erstellt eine leere Warteschlange
func NewQueue() *Queue {
	return &Queue{elements: []string{}}
}

//Enqueue fügt ein Element hinten an die Warteschlange an
func (q *Queue) Enqueue(element string) {
	q.elements = append(q.elements, element)
}

//Dequeue nimmt das vorderste Element weg und gibt es zurück
//Gibt zusätzlich ein 'bool' zurück, das falsch ist, wenn die Queue leer war
func (q *Queue) Dequeue() (string, bool) {
	if q.IsEmpty() {
		return "", false
	}
	//Das erste Element herausholen
	element := q.elements[0]
	//Das erste Element aus dem Slice löschen
	q.elements = q.elements[1:]
	return element, true
}

//IsEmpty prüft, ob die Warteschlange leer ist
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}
