package main

import (
	"fmt"
)

func main() {
	arr1 := []int{2, 1}       // f
	arr2 := []int{3, 5, 5}    // f
	arr3 := []int{0, 3, 2, 1} // t
	arr4 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(validMountainArray(arr1))
	fmt.Println(validMountainArray(arr2))
	fmt.Println(validMountainArray(arr3))
	fmt.Println(validMountainArray(arr4))
}

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	i, j := 0, len(A)-1
	for i < len(A)-1 && A[i] < A[i+1] {
		i++
	}
	for j > 0 && A[j] < A[j-1] {
		j--
	}
	return i > 0 && j < len(A)-1 && i == j
}
