package main

import "fmt"

func main() {
	printPrimeFactor(84)
}

func printPrimeFactor(n int) {
	if n <= 1 {
		return
	}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			fmt.Println(i)
			n = n / i
		}
	}
	if n > 1 {
		fmt.Println(n)
	}
}
