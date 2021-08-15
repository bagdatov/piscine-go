package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	divider := ""
	for _, value := range s {
		if value != ' ' && value != '\n' && value != '	' {
			divider += string(value)
		} else {
			if divider != "" {
				result = append(result, divider)
				divider = ""
			}
		}
	}
	result = append(result, divider)
	return result
}
