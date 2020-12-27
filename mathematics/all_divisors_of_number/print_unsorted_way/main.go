package main

import "fmt"

func main() {
	printAllDivisors(25)
}

func printAllDivisors(n int) {
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
			if i != n/i {
				fmt.Println(n / i)
			}
		}
	}
}
