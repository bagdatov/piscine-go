package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	for l.Head != nil {
		if l.Head.Data == data_ref {
			l.Head = l.Head.Next
		} else {
			break
		}
	}
	element := l.Head
	prev := element
	for element != nil {
		if element.Next != nil {
			if element.Data != data_ref {
				prev = element
			} else {
				prev.Next = element.Next
			}
			element = element.Next
		} else {
			if element.Data == data_ref {
				prev.Next = element.Next
			}
			element = element.Next
		}
	}
}
