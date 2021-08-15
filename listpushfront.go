package piscine

func ListPushFront(l *List, data interface{}) {
	element := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = element
		l.Tail = element
		return
	}
	element.Next = l.Head
	l.Head = element
}
