package graph

import "testing"

func TestMinHeapSorting(t *testing.T) {
	h := NewMinHeap()

	//Elemente unsortiert hinzufügen
	h.Push(HeapItem{Id: "C", Distance: 42.5})
	h.Push(HeapItem{Id: "A", Distance: 10.2})
	h.Push(HeapItem{Id: "D", Distance: 100.0})
	h.Push(HeapItem{Id: "B", Distance: 5.1})

	//das kleinste Element muss jetzt "B" mit 5.1 sein
	item, ok := h.Pop()
	if !ok || item.Id != "B" || item.Distance != 5.1 {
		t.Fatalf("Erwartet: B (5.1), bekommen: %v (%v)", item.Id, item.Distance)
	}

	//danach muss "A" mit 10.2 kommen
	item, ok = h.Pop()
	if !ok || item.Id != "A" || item.Distance != 10.2 {
		t.Fatalf("Erwartet: A (10.2), bekommen: %v (%v)", item.Id, item.Distance)
	}

	//danach "C" mit 42.5
	item, ok = h.Pop()
	if !ok || item.Id != "C" || item.Distance != 42.5 {
		t.Fatalf("Erwartet: C (42.5), bekommen: %v (%v)", item.Id, item.Distance)
	}

	//als letztes das größte, also "D" mit 100.0
	item, ok = h.Pop()
	if !ok || item.Id != "D" || item.Distance != 100.0 {
		t.Fatalf("Erwartet: D (100.0), bekommen: %v (%v)", item.Id, item.Distance)
	}

	//Heap müsste jetzt leer sein
	if !h.IsEmpty() {
		t.Fatalf("Heap sollte jetzt leer sein!")
	}
}
