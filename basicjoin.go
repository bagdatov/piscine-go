package piscine

func BasicJoin(elems []string) string {
	asnwer := ""
	for i := 0; i < len(elems); i++ {
		asnwer += elems[i]
	}
	return asnwer
}
