package piscine

func ListClear(l *List) {
	elementToErase := l.Head
	for elementToErase != nil {
		elementToErase.Data = nil
		elementToErase = elementToErase.Next
	}
	l.Head = nil
}
