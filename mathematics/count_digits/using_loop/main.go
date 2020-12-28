package main

import "fmt"

func main() {
	d := digitCount(99999)
	fmt.Println(d)
}

func digitCount(n int) int {
	result := 0
	for n > 0 {
		result++
		n = n / 10
	}
	return result
}
