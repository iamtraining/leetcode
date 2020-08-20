package main

import "fmt"

func main() {
	arr := []int{-4, -1, 0, 3, 10}
	arr1 := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(arr))
	fmt.Println(sortedSquares(arr1))
}

func sortedSquares(A []int) []int {
	lenA := len(A)
	for _, num := range A {
		A = append(A, num*num)
	}
	return quicksort(A[lenA:])
}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		basis, less, greater := arr[0], []int{}, []int{}
		for _, num := range arr[1:] {
			if basis > num {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}
		}
		less = quicksort(less)
		less = append(less, basis)
		greater = quicksort(greater)
		return append(less, greater...)
	}
}
