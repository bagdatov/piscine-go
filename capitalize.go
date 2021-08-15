package piscine

func Capitalize(s string) string {
	capWord := true
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' || s[i] >= '0' && s[i] <= '9' {
			if capWord {
				result += ToUpper(string(s[i]))
				capWord = false
			} else {
				result += ToLower(string(s[i]))
			}
		} else {
			capWord = true
			result += string(s[i])
		}
	}
	return result
}
