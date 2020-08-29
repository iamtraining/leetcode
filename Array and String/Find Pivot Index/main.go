package main

import "fmt"

func main() {
	arr1 := []int{1, 7, 3, 6, 5, 6}
	arr2 := []int{1, 2, 3}
	fmt.Println(pivotIndex(arr1))
	fmt.Println(pivotIndex(arr2))
}

func pivotIndex(nums []int) int {
	sum, sumleft := 0, 0
	for _, num := range nums {
		sum += num
	}
	for i := 0; i < len(nums); i++ {
		if sumleft == sum-sumleft-nums[i] {
			return i
		}
		sumleft += nums[i]
	}
	return -1
}
