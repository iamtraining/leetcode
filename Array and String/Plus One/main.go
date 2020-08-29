package main

import "fmt"

func main() {
	arr1 := []int{3, 6, 1, 0}
	arr2 := []int{1, 2, 3, 4}
	arr3 := []int{0, 0, 1, 2}
	arr4 := []int{0, 0, 0, 1}
	arr5 := []int{0, 0, 0, 9}
	arr6 := []int{0, 0, 9, 9}
	arr7 := []int{0, 9, 9, 9}
	arr8 := []int{9, 9, 9, 9}
	arr9 := []int{6, 8, 9, 9}
	fmt.Println(plusOne(arr1))
	fmt.Println(plusOne(arr2))
	fmt.Println(plusOne(arr3))
	fmt.Println(plusOne(arr4))
	fmt.Println(plusOne(arr5))
	fmt.Println(plusOne(arr6))
	fmt.Println(plusOne(arr7))
	fmt.Println(plusOne(arr8))
	fmt.Println(plusOne(arr9))
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits = append([]int{1}, digits...)
	return digits
}
