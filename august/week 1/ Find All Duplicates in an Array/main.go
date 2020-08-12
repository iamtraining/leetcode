package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(arr))
}

func findDuplicates(nums []int) []int {
	checkmap := make(map[int]bool)
	var dups []int
	for _, arg := range nums {
		if _, ok := checkmap[arg]; ok {
			dups = append(dups, arg)
		} else {
			checkmap[arg] = true
		}
	}
	return dups
}
