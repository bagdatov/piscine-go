package piscine

func Index(s string, toFind string) int {
	length := len(toFind)
	for i := 0; i < len(s); i++ {
		if length+i <= len(s) {
			if s[i:length+i] == toFind {
				return i
			}
		}
	}
	return -1
}
