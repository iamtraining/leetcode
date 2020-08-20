package main

import "fmt"

func main() {
	arr1 := []int{1, 1, 2}
	arr2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr1))
	fmt.Println(removeDuplicates(arr2))
}

func removeDuplicates(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}
