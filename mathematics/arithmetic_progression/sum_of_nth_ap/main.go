package main

import "fmt"

func main() {
	s := helperGetAP(4, 3, 10)
	fmt.Println(s)

	sum := getSumOfNthAP(s, 5)
	fmt.Println(sum)
}

func getSumOfNthAP(s []int, n int) int {
	first := s[0]
	d := s[1] - s[0]

	sum := (float64(n) / 2) * float64(2*first+(n-1)*d)
	return int(sum)
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
