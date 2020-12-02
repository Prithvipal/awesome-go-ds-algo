package main

import "fmt"

func main() {
	arr := []int{64, 25, 12, 22, 11}

	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		minVal := arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = minVal
	}
	fmt.Println(arr)
}
