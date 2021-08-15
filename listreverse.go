package piscine

func ListReverse(l *List) {
	linkReversed := &List{}
	traverse := l.Head
	for traverse != nil {
		ListPushFront(linkReversed, traverse.Data)
		traverse = traverse.Next
	}
	l.Head = linkReversed.Head
	l.Tail = linkReversed.Tail
}
