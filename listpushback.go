package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	element := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = element
	} else {
		l.Tail.Next = element
	}
	l.Tail = element
}
