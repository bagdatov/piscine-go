package piscine

func ListSize(l *List) int {
	counter := 0
	element := l.Head
	for element != nil {
		counter++
		element = element.Next
	}
	return counter
}
