package main

import "fmt"

func main() {
	fmt.Println(minSubArrayLen1(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen2(7, []int{2, 3, 1, 2, 4, 3}))

}

func minSubArrayLen1(s int, nums []int) int {
	n, ans := len(nums), 0
	if n == 0 {
		return ans
	}
	sums := []int{nums[0]}
	for i := 1; i < n; i++ {
		sums = append(sums, sums[i-1]+nums[i])
	}

	for i := 0; i < len(sums); i++ {
		if sums[i] >= s && (ans == 0 || ans > 1) {
			ans = i + 1
		}
		for j := i + 1; j < len(sums); j++ {
			if sums[j]-sums[i] >= s && (ans == 0 || ans > j-i) {
				ans = j - i
				break
			}
		}
	}

	return ans
}

func minSubArrayLen2(s int, nums []int) int {
	var i, j, sum int
	min := 1<<31 - 1
	for j < len(nums) {
		sum += nums[j]
		j++
		for sum >= s {
			if min > j-i {
				min = j - i
			}
			sum -= nums[i]
			i++
		}
	}
	if min == 1<<31-1 {
		return 0
	}
	return min
}
