package main

import "fmt"

func main() {
	arr1 := []int{0, 1, 0, 3, 12}
	moveZeroes(arr1)
	fmt.Println(arr1)
}

func moveZeroes(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}
	for i, zeroes := len(nums)-1, len(nums)-count; i >= 0 && zeroes > 0; i-- {
		nums[i] = 0
		zeroes--
	}
}
