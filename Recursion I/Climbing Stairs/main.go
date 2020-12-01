package main

import "fmt"

func main() {
	fmt.Println(climbStairs1(2))
	fmt.Println(climbStairs1(3))
	fmt.Println(climbStairs2(2))
	fmt.Println(climbStairs2(3))
	fmt.Println(climbStairs3(2))
	fmt.Println(climbStairs3(3))
}

// brute force
func climbStairs1(n int) int {
	return helper1(0, n)
}

func helper1(res int, n int) int {
	if res > n {
		return 0
	}
	if res == n {
		return 1
	}
	return helper1(res+1, n) + helper1(res+2, n)
}

// recursion with memoization
func climbStairs2(n int) int {
	cache := make([]int, n)
	return helper2(cache, n, 0)
}

func helper2(arr []int, n int, res int) int {
	if res > n {
		return 0
	}
	if res == n {
		return 1
	}
	if arr[res] > 0 {
		return arr[res]
	}
	arr[res] = helper2(arr, n, res+1) + helper2(arr, n, res+2)
	return arr[res]
}

// fibonacci number
func climbStairs3(n int) int {
	if n == 1 {
		return 1
	}
	first, second := 1, 2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}
