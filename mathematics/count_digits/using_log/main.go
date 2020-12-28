package main

import (
	"fmt"
	"math"
)

func main() {
	n := 99999
	d := int(math.Log10(float64(n))) + 1
	fmt.Println(d)
}
