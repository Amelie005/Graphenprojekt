package graph

//stellt ein Item in dem Heap dar
type HeapItem struct {
	Id       string
	Distance float64
}

//minHeap mit Elementen
type MinHeap struct {
	elements []HeapItem
}

//neuer minHeap
func NewMinHeap() *MinHeap {
	return &MinHeap{elements: []HeapItem{}}
}

//gibt an ob minHeap leer ist
func (h *MinHeap) IsEmpty() bool {
	return len(h.elements) == 0
}

//fügt ein neues Element am Ende hinzu und lässt es nach oben wandern
func (h *MinHeap) Push(item HeapItem) {
	h.elements = append(h.elements, item)
	h.upHeap(len(h.elements) - 1)
}

//holt das kleinste Element von der Spitze (Index 0)
func (h *MinHeap) Pop() (HeapItem, bool) {
	if h.IsEmpty() {
		return HeapItem{}, false
	}

	//kleinstes Element ist immer ganz vorne
	root := h.elements[0]

	//letztes Element an die Spitze setzen
	lastIndex := len(h.elements) - 1
	h.elements[0] = h.elements[lastIndex]
	h.elements = h.elements[:lastIndex] //letztes Element entfernen

	//neue Spitze nach unten sinken lassen, bis sie richtig sitzt
	if !h.IsEmpty() {
		h.downHeap(0)
	}

	return root, true
}

//schiebt ein Element nach oben, wenn es kleiner als sein Vater ist
func (h *MinHeap) upHeap(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		//wenn das aktuelle Element kleiner ist als sein Vater, dann tauschen
		if h.elements[index].Distance < h.elements[parentIndex].Distance {
			h.elements[index], h.elements[parentIndex] = h.elements[parentIndex], h.elements[index]
			index = parentIndex
		} else {
			break
		}
	}
}

//lässt ein Element nach unten sinken, wenn seine Kinder kleiner sind
func (h *MinHeap) downHeap(index int) {
	lastIndex := len(h.elements) - 1

	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		//prüfen, ob das linke Kind kleiner ist
		if leftChild <= lastIndex && h.elements[leftChild].Distance < h.elements[smallest].Distance {
			smallest = leftChild
		}
		//prüfen, ob das rechte Kind noch kleiner ist
		if rightChild <= lastIndex && h.elements[rightChild].Distance < h.elements[smallest].Distance {
			smallest = rightChild
		}

		//wenn das kleinste Element nicht mehr das aktuelle ist, tauschen und weiter absinken
		if smallest != index {
			h.elements[index], h.elements[smallest] = h.elements[smallest], h.elements[index]
			index = smallest
		} else {
			break
		}
	}
}
