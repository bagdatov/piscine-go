package main

import "fmt"

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func pushFront(node *NodeAddL, num int) *NodeAddL {
	var tmp NodeAddL
	tmp.Num = num
	tmp.Next = node
	return &tmp
}

func AddLinkedNumbers(num1, num2 *NodeAddL) *NodeAddL {
	l := &NodeAddL{Num: 0}
	k := l
	tmp := 0
	prev := l
	for num1 != nil && num2 != nil {
		if num1 != nil {
			tmp += num1.Num
			num1 = num1.Next
		}
		if num2 != nil {
			tmp += num2.Num
			num2 = num2.Next
		}
		if num1 != nil || num2 != nil {
			if tmp > 9 {
				div := tmp / 10
				prev.Num += div
				tmp = tmp % 10
			}
			l.Num = tmp
			tmp = 0
			l.Next = &NodeAddL{Num: 0}
			prev = l
			l = l.Next
		}
	}
	if tmp > 9 {
		div := tmp / 10
		prev.Num += div
		tmp = tmp % 10
	}
	l.Num = tmp
	return k
}

func main() {
	// 6 -> 7 -> 8 -> 5
	num1 := &NodeAddL{Num: 5}
	num1 = pushFront(num1, 8)
	num1 = pushFront(num1, 7)
	num1 = pushFront(num1, 6)

	// 9 -> 9 -> 9 -> 2
	num2 := &NodeAddL{Num: 2}
	num2 = pushFront(num2, 9)
	num2 = pushFront(num2, 9)
	num2 = pushFront(num2, 9)

	// 1 -> 6 -> 7 -> 7 -> 7
	result := AddLinkedNumbers(num1, num2)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
