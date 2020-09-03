package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(1000))
	fmt.Println(numSquares(2))
	fmt.Println(numSquares(3))
	fmt.Println(numSquares(4))
	fmt.Println(numSquares(5))
	fmt.Println(numSquares(6))
	fmt.Println(numSquares(7))

}

func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
