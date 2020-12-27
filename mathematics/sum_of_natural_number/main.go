package main

import "fmt"

func main() {
	fmt.Println(sum(3))
}

func sum(n int) int {
	return (n * (n + 1)) / 2
}
