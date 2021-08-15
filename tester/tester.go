package main

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	element := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = element
	} else {
		element.Next = l.Head
	}
	l.Head = element
}

func main() {

	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "man")
	ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}
