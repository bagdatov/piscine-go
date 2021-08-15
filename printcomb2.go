package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for q := '0'; q <= '9'; q++ {
				for k := '0'; k <= '9'; k++ {
					if i == '9' && j == '8' && q == '9' && k == '9' {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(q)
						z01.PrintRune(k)
					} else if q >= i && k > j {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(q)
						z01.PrintRune(k)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else if q > i && k >= j {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(q)
						z01.PrintRune(k)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else if q > i && k < j {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(q)
						z01.PrintRune(k)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
