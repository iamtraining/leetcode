package main

import "fmt"

func main() {
	fmt.Println(fib(10))
}

func fib(N int) int {
	cache := map[int]int{}
	cache[0], cache[1] = 0, 1
	return helper(N, cache)
}

func helper(N int, m map[int]int) int {
	if num, ok := m[N]; ok {
		return num
	}
	m[N] = helper(N-1, m) + helper(N-2, m)
	return m[N]
}
