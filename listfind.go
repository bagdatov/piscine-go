package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	element := l.Head
	for element != nil {
		if comp(element.Data, ref) {
			return &element.Data
		}
		element = element.Next
	}
	return nil
}
