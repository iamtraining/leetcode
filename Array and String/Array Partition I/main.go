package main

import "fmt"

func main() {
	arr := []int{1, 4, 3, 2}
	fmt.Println(arrayPairSum(arr))
}

func arrayPairSum(nums []int) int {
	nums = quickSort(nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		i++
	}
	return sum
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	} else {
		basic := nums[0]
		less, greater := []int{}, []int{}
		for _, num := range nums[1:] {
			if num < basic {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}

		}
		less = quickSort(less)
		less = append(less, basic)
		greater = quickSort(greater)
		return append(less, greater...)
	}

}
