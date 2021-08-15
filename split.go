package piscine

func Split(s, sep string) []string {
	var result []string
	n := 0
	for i := 0; i < len(s)-len(sep); {
		for j := i; j < len(s)-len(sep); j++ {
			if s[j:j+len(sep)] == sep {
				result = append(result, s[i:j])
				n = j
				i = j + len(sep)
			}
		}
		i++
	}
	result = append(result, s[n+len(sep):])
	return result
}
