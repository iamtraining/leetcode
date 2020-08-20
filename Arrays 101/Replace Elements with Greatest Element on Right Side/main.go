package main

import (
	"fmt"
)

func main() {
	arr1 := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(arr1)) // 18,6,6,6,1,-1
}

func replaceElements(arr []int) []int {
	l := len(arr)
	check := arr[l-1]
	arr[l-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		t := arr[i]
		arr[i] = check
		if t > check {
			check = t
		}
	}
	return arr
}
