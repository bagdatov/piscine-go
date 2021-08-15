package piscine

func NRune(s string, n int) rune {
	var runeN rune = 0
	if len(s) >= n && n >= 1 {
		runeN = ([]rune(s))[n-1]
	}
	return runeN
}
