package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func main() {

	g1 := gcd(50, 125)
	fmt.Println(g1)
	g2 := gcd(125, 50)
	fmt.Println(g2)
}
