package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(arr)
	storeCombination(arr, r, n)
	// the method was initially mentioned at
	// https://www.geeksforgeeks.org/print-all-possible-combinations-of-r-elements-in-a-given-array-of-size-n/
}

func storeCombination(arr []int, r, n int) {
	// A temporary array to store
	// all combination one by one
	data := make([]int, n)
	combinationUtil(arr, r, n, 0, data, 0)
}

func combinationUtil(arr []int, r, n, index int, data []int, i int) {
	if index == n {
		// current combination matches, so we print it
		if lastElement(data) {
			// but first we check if this is a last possilbe element
			for _, v := range data {
				z01.PrintRune(rune(v + '0'))
			}
			z01.PrintRune('\n')
			return
		} else {
			// if not last we print it
			for j := 0; j < n; j++ {
				z01.PrintRune(rune(data[j] + '0'))
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
			return
		}
	}
	// When no more elements are
	// there to put in data[]
	if i >= r {
		return
	}
	// current is included, put
	// next at next location
	data[index] = arr[i]
	combinationUtil(arr, r, n, index+1, data, i+1)
	// current is excluded, replace it
	// with next (Note that i+1 is passed,
	// but index is not changed)
	combinationUtil(arr, r, n, index, data, i+1)
}

func lastElement(data []int) bool {
	// hardcoded last possible element if 0 < n < 10
	if areSlicesEqual(data, []int{9}) {
		return true
	} else if areSlicesEqual(data, []int{8, 9}) {
		return true
	} else if areSlicesEqual(data, []int{7, 8, 9}) {
		return true
	} else if areSlicesEqual(data, []int{6, 7, 8, 9}) {
		return true
	} else if areSlicesEqual(data, []int{5, 6, 7, 8, 9}) {
		return true
	} else if areSlicesEqual(data, []int{4, 5, 6, 7, 8, 9}) {
		return true
	} else if areSlicesEqual(data, []int{3, 4, 5, 6, 7, 8, 9}) {
		return true
	} else if areSlicesEqual(data, []int{2, 3, 4, 5, 6, 7, 8, 9}) {
		return true
	} else if areSlicesEqual(data, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		return true
	}
	return false
}

func areSlicesEqual(a []int, b []int) bool {
	// compare slices
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
