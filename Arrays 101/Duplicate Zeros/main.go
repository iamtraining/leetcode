package main

import "fmt"

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	arr1 := []int{1, 2, 3}
	duplicateZeros(arr)
	duplicateZeros(arr1)
	fmt.Println(arr)
	fmt.Println(arr1)

}

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 && i+1 < len(arr) {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			i++
		}
	}
}
