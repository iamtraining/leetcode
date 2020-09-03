package main

import "fmt"

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, S int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	res := 0
	helper(nums, S, 0, 0, &res)
	return res
}

func helper(nums []int, S int, ind int, sum int, res *int) {
	fmt.Println(ind, "line 19")
	fmt.Println(sum+nums[ind], "line 22")
	sum += nums[ind]
	if len(nums)-1 == ind {
		if S == sum {
			fmt.Println(*res+1, "line 26")
			*res = *res + 1
		}
	} else {
		helper(nums, S, ind+1, sum, res)
	}
	fmt.Println(sum-nums[ind], "line 32")
	sum -= nums[ind]

	fmt.Println(sum-nums[ind], "line 35")
	sum -= nums[ind]
	if len(nums)-1 == ind {
		if S == sum {
			fmt.Println(*res+1, "line 39")
			*res = *res + 1
		}
	} else {
		helper(nums, S, ind+1, sum, res)
	}
}
