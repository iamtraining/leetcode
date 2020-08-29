package main

import "fmt"

func main() {
	arr1 := []int{3, 6, 1, 0}
	arr2 := []int{1, 2, 3, 4}
	arr3 := []int{0, 0, 1, 2}
	arr4 := []int{0, 0, 0, 1}
	fmt.Println(dominantIndex(arr1))
	fmt.Println(dominantIndex(arr2))
	fmt.Println(dominantIndex(arr3))
	fmt.Println(dominantIndex(arr4))
}

func dominantIndex(nums []int) int {
	first, second, ind := 0, 0, 0
	for i, num := range nums {
		if num > first {
			second, first, ind = first, num, i
		} else if num > second {
			second = num
		}
	}
	if first >= second*2 {
		return ind
	} else {
		return -1
	}
}
