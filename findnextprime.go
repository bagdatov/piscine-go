package piscine

func FindNextPrime(nb int) int {
	if nb > 1 {
		if IsPrime(nb) {
			return nb
		} else {
			for i := nb + 1; i < nb*2; i++ {
				if IsPrime(i) {
					return i
				}
			}
		}
	}
	return 2
}
