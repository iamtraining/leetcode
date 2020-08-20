package main

import "fmt"

func main() {
	arr1 := []int{3, 2, 2, 3}             // val 3, len 2
	arr2 := []int{0, 1, 2, 2, 3, 0, 4, 2} // val 2, len 5
	fmt.Println(removeElement(arr1, 3))
	fmt.Println(removeElement(arr2, 2))
}

func removeElement(nums []int, val int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}
