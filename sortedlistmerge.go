package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	result := &NodeI{}
	store := result
	for n1 != nil && n2 != nil {
		if n1.Data > n2.Data {
			store.Next = &NodeI{Data: n2.Data}
			store = store.Next
			n2 = n2.Next
		} else {
			store.Next = &NodeI{Data: n1.Data}
			store = store.Next
			n1 = n1.Next
		}
		if n1 == nil {
			store.Next = &NodeI{Data: n2.Data}
		} else if n2 == nil {
			store.Next = &NodeI{Data: n1.Data}
		}
	}
	return result.Next
}
