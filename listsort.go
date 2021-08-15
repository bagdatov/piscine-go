package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	k := l
	for k != nil {
		j := k.Next
		for j != nil {
			if k.Data > j.Data {
				k.Data, j.Data = j.Data, k.Data
			}
			j = j.Next
		}
		k = k.Next
	}
	return l
}
