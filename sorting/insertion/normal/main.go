package main

import "fmt"

func main() {
	arr := []int{12, 11, 13, 5, 6}

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	fmt.Println(arr)
}
