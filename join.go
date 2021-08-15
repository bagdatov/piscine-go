package piscine

func Join(strs []string, sep string) string {
	answer := ""
	for i := 0; i < len(strs); i++ {
		answer += strs[i]
		if i != len(strs)-1 {
			answer += sep
		}
	}
	return answer
}
