package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPowerOfFour(1))
	fmt.Println(isPowerOfFour(5))
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(-2147483648))
}

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	} else {
		ok := math.Log(float64(num)) / math.Log(4)
		return math.Floor(ok) == ok
	}
}
