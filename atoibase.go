package piscine

func AtoiBase(s string, base string) int {
	var sArray []string
	var baseArray []string
	result := 0
	power := 0
	for _, value := range base {
		baseArray = append(baseArray, string(value))
	}
	for _, value := range s {
		sArray = append(sArray, string(value))
	}
	if IsValidBase(s, base) {
		for i := len(sArray) - 1; i >= 0; i-- {
			for j := 0; j < len(base); j++ {
				if sArray[i] == baseArray[j] {
					result = j*(RecursivePower(len(base), power)) + result
					power++
				}
			}
		}
		return result
	}
	return 0
}

func IsValidBase(s string, base string) bool {
	if len(base) >= 2 {
		for i := 0; i < len(base)-1; i++ {
			if base[i] != '+' && base[i] != '-' {
				for j := i + 1; j < len(base); j++ {
					if base[i] == base[j] {
						return false
					}
				}
			} else {
				return false
			}
		}
		return true
	}
	return false
}
