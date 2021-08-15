package main

import "fmt"

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func Changeorder(node *NodeAddL) *NodeAddL {
	var copy *NodeAddL
	it := node
	for it != nil {
		if it.Num%2 == 1 {
			copy = pushBack(copy, it.Num)
		}
		it = it.Next
	}
	it = node
	for it != nil {
		if it.Num%2 == 0 {
			copy = pushBack(copy, it.Num)
		}
		it = it.Next
	}
	return copy
}

func pushBack(node *NodeAddL, num int) *NodeAddL {
	tmp := &NodeAddL{Num: num}
	if node == nil {
		return tmp
	}
	it := node
	for it.Next != nil {
		it = it.Next
	}
	it.Next = tmp
	return node
}

func main() {
	num1 := &NodeAddL{Num: 1}
	num1 = pushBack(num1, 2)
	num1 = pushBack(num1, 3)
	num1 = pushBack(num1, 4)
	num1 = pushBack(num1, 5)

	result := Changeorder(num1)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
