package main

import "fmt"

func main() {
	arr1 := []int{2, 7, 11, 15}
	arr2 := []int{3, 2, 4}
	arr3 := []int{3, 3}
	fmt.Println(twoSum(arr1, 9))
	fmt.Println(twoSum(arr2, 6))
	fmt.Println(twoSum(arr3, 6))
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	result := []int{}
	for i, v := range nums {
		if res, ok := m[target-v]; ok {
			result = append(result, res, i)
		}
		m[v] = i
	}
	return result
}
