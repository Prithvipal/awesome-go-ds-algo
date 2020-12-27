package main

import "fmt"

func main() {
	n := 101
	primes := getPrimes(n)
	for i := 2; i <= n; i++ {
		if primes[i] {
			fmt.Println(i)
			for j := i * i; j <= n; j = j + i {
				primes[j] = false
			}
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
