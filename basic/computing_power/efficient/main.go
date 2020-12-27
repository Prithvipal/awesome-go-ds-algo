package main

import "fmt"

func main() {
	p := power(4, 3)
	fmt.Println(p)
}

func power(x, n int) int {
	if n == 0 {
		return 1
	}
	temp := power(x, n/2)
	temp = temp * temp
	if n%2 == 0 {
		return temp
	}
	return temp * x

}
