package main

import "fmt"

func main() {
	arr := []int{1, 1, 0, 1, 1, 1}
	arr1 := []int{1, 0, 1, 1, 0, 1}
	fmt.Println(findMaxConsecutiveOnes(arr))
	fmt.Println(findMaxConsecutiveOnes(arr1))
}

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	max := 0

	for _, v := range nums {
		if v == 1 {
			count++

		}
		if count > max {
			max = count
		}
		if v == 0 {
			count = 0
		}
	}
	return max
}

func findMaxConsecutiveOnes1(nums []int) int {
	temp, max := 0, 0
	for _, n := range nums {
		if n == 1 {
			temp++
		}
		if temp > max {
			max = temp
		}
		if n == 0 {
			temp = 0
		}
	}
	return max
}
