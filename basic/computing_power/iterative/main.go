package main

import "fmt"

func main() {
	p := power(5, 2)
	fmt.Println(p)
}

func power(x, n int) int {
	result := 1
	for n > 0 {
		if n%2 != 0 {
			result = result * x
		}
		x = x * x
		n = n / 2
	}
	return result
}
