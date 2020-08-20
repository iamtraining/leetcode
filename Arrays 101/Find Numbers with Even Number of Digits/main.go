package main

import "fmt"

func main() {
	arr := []int{555, 901, 482, 1771}  // 1
	arr1 := []int{12, 345, 2, 6, 7896} // 2
	fmt.Println(findNumbers(arr))
	fmt.Println(findNumbers(arr1))
}

func findNumbers(nums []int) int {
	n, even := 0, 0
	for _, v := range nums {
		for v > 0 {
			v /= 10
			n++
		}
		if n%2 == 0 {
			even++
		}
		n = 0
	}
	return even
}
