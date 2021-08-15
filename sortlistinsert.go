package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	k := l
	if k.Data > data_ref {
		newNode := &NodeI{k.Data, k.Next}
		currentNode := &NodeI{data_ref, newNode}
		k = currentNode
		return k
	}
	oldL := l
	currentL := l.Next
	for currentL != nil {
		if currentL.Data > data_ref {
			newNode := &NodeI{data_ref, currentL}
			oldL.Next = newNode
			return l
		}
		currentL = currentL.Next
		oldL = oldL.Next
	}
	newNode := &NodeI{data_ref, nil}
	oldL.Next = newNode
	return l
}
