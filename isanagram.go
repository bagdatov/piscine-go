package piscine

func IsAnagram(str1, str2 string) bool {
	a := []rune{}
	for _, v := range str1 {
		if v != ' ' {
			a = append(a, v)
		}
	}
	b := []rune{}
	for _, v := range str2 {
		if v != ' ' {
			b = append(b, v)
		}
	}
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}

		}
	}
	for i := 0; i < len(b)-1; i++ {
		for j := 0; j < len(b)-i-1; j++ {
			if b[j] > b[j+1] {
				b[j], b[j+1] = b[j+1], b[j]
			}

		}
	}
	str1 = string(a)
	str2 = string(b)
	return str1 == str2
}
