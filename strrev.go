package piscine

func StrRev(s string) string {
	reversed := ""
	for _, a := range s {
		reversed = string(a) + reversed
	}
	return reversed
}
