package main

import "fmt"

func main() {
	printAllDivisors(15)
}

func printAllDivisors(n int) {
	i := 1
	for ; i*i < n; i++ {
		if n%i == 0 {
			fmt.Println(i)
		}
	}

	for ; i >= 1; i-- {
		if n%i == 0 {
			fmt.Println(n / i)
		}
	}
}
