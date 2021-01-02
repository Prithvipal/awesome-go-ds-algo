package main

import "fmt"

func main() {
	var a uint32 = 5
	x := ^a
	fmt.Println(x)
	var b int32 = 5
	y := ^b
	fmt.Println(y)
}
