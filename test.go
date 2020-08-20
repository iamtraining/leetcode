package main

import "fmt"

func main() {
	arr1 := []int{-4, -1, 0, 3, 10}
	arr2 := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(arr1))
	fmt.Println(sortedSquares(arr2))
}

func sortedSquares(A []int) []int {
	lA := len(A)
	for _, num := range A {
		A = append(A, num*num)
	}
	return quicksort(A[lA:])
}

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	} else {
		basis, less, greater := nums[0], []int{}, []int{}
		for _, num := range nums[1:] {
			if num < basis {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}

		}
		less = quicksort(less)
		less = append(less, basis)
		greater = quicksort(greater)
		return append(less, greater...)
	}
}
