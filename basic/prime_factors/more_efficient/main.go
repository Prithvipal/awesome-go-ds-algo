package main

import "fmt"

func main() {
	printPrimeFactor(84)
}

func printPrimeFactor(n int) {
	if n <= 1 {
		return
	}
	for n%2 == 0 {
		fmt.Println(2)
		n = n / 2
	}
	for n%3 == 0 {
		fmt.Println(3)
		n = n / 3
	}
	for i := 5; i*i <= n; i = i + 6 {
		for n%i == 0 {
			fmt.Println(i)
			n = n / i
		}
		for n%(i+2) == 0 {
			fmt.Println(i + 2)
			n = n / (n + 2)
		}
	}
	if n > 3 {
		fmt.Println(n)
	}
}
