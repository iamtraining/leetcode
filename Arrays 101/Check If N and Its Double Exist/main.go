package main

import "fmt"

func main() {
	arr1 := []int{10, 2, 5, 3}
	arr2 := []int{7, 1, 14, 11}
	arr3 := []int{3, 1, 7, 11}
	//arr4 := []int{}
	fmt.Println(checkIfExist(arr1))
	fmt.Println(checkIfExist(arr2))
	fmt.Println(checkIfExist(arr3))
}

func checkIfExist(arr []int) bool {
	m := make(map[int]bool)
	for _, num := range arr {

		if _, ok := m[num*2]; ok {
			return true
		}
		if num%2 == 0 {
			if _, ok := m[num/2]; ok {
				return true
			}
		}
		m[num] = true
	}
	return false
}
