package main

import "fmt"

func main() {
	n := 101
	isPrime := getPrimes(n)
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * 2; j <= n; j = j + i {
				isPrime[j] = false
			}
		}
	}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			fmt.Println(i)
		}
	}
}

func getPrimes(n int) []bool {
	s := make([]bool, n+1)
	for i := range s {
		s[i] = true
	}
	return s
}
