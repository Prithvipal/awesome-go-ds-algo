package main

import "fmt"

func main() {
	l := lcm(12, 15)
	fmt.Println(l)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
