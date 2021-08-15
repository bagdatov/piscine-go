package piscine

func ListMerge(l1 *List, l2 *List) {
	element := l2.Head
	for element != nil {
		ListPushBack(l1, element.Data)
		element = element.Next
	}
}
