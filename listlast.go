package piscine

func ListLast(l *List) interface{} {
	element := l.Tail
	var answer interface{}
	for element != nil {
		answer = l.Tail.Data
		element = element.Next
	}
	return answer
}
