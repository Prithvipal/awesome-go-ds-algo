package main

import "fmt"

func main() {
	s := helperGetAP(4, 3, 10)
	fmt.Println(s)
	nthAP := getNthAP(s, 10)
	fmt.Println(nthAP)
}

// Actual logic to find nth AP.
func getNthAP(s []int, n int) int {
	first := s[0]
	d := s[1] - s[0]
	nthAP := first + (n-1)*d
	return nthAP
}

// a simple helper function to return a slice for AP.
// It is not part of logic. It just provide slice.
func helperGetAP(n, d, size int) []int {
	s := make([]int, size)
	s[0] = n
	for i := 1; i < size; i++ {
		s[i] = s[i-1] + d
	}
	return s
}
