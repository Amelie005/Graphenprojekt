package graph

import "testing"

func TestQueueFIFO(t *testing.T) {
	q := NewQueue()

	//prüfen, ob eine neue Queue tatsächlich leer ist
	if !q.IsEmpty() {
		t.Fatalf("Neue Queue sollte leer sein!")
	}

	//drei Elemente anstellen
	q.Enqueue("Erster")
	q.Enqueue("Zweiter")
	q.Enqueue("Dritter")

	if q.IsEmpty() {
		t.Fatalf("Queue sollte nicht mehr leer sein!")
	}

	//Elemente herausholen und prüfen, ob die Reihenfolge stimmt (also FIFO)
	val, ok := q.Dequeue()
	if !ok || val != "Erster" {
		t.Fatalf("Erwartet: 'Erster', Got: '%v'", val)
	}

	val, ok = q.Dequeue()
	if !ok || val != "Zweiter" {
		t.Fatalf("Erwartet: 'Zweiter', Got: '%v'", val)
	}

	val, ok = q.Dequeue()
	if !ok || val != "Dritter" {
		t.Fatalf("Erwartet: 'Dritter', Got: '%v'", val)
	}

	//Queue muss wieder leer sein
	if !q.IsEmpty() {
		t.Fatalf("Queue sollte nach dem Leeren leer sein!")
	}

	//Dequeue bei leerer Queue darf nicht abstürzen
	_, ok = q.Dequeue()
	if ok {
		t.Fatalf("Dequeue bei leerer Queue sollte false zurückgeben!")
	}
}
