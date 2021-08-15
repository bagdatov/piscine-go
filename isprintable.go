package piscine

func IsPrintable(s string) bool {
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= rune(0) && rune(s[i]) <= rune(31) {
			return false
		}
	}
	return true
}
