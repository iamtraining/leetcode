package main

import "fmt"

func main() {
	arr := []int{3, 1, 2, 4} // 2,4,3,1 / , 2,4,1,3 / , /4,2,1,3
	fmt.Println(sortArrayByParity(arr))
}

func sortArrayByParity(A []int) []int {

	for i, j := 0, len(A)-1; i < j; {
		if A[i]%2 > A[j]%2 {
			tmp := A[i]
			A[i] = A[j]
			A[j] = tmp

		}
		if A[i]%2 == 0 {
			i++
		}
		if A[j]%2 == 1 {
			j--
		}
	}
	return A
}
