package main

import "fmt"

func main() {
	n := 260
	count1 := getTrailingZeroInFact(n)
	fmt.Println(count1)
	count2 := countTrailignZeros(n)
	fmt.Println(count2)

}

// using while loop
func getTrailingZeroInFact(n int) int {
	d := 5
	count := 0
	for n/d != 0 {
		count = count + (n / d)
		d = d * 5
	}
	return count
}

func countTrailignZeros(n int) int {
	count := 0
	for i := 5; i < n; i = i * 5 {
		count = count + (n / i)
	}
	return count
}
