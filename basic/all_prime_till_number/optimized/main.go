package main

import "fmt"

func main() {
	n := 100
	primes := getPrimes(n)
	for i := 2; i*i <= n; i++ {
		if primes[i] {
			for j := i * i; j <= n; j = j + i {
				primes[j] = false
			}
		}
	}
	for i := 2; i <= n; i++ {
		if primes[i] {
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
