package main

import "fmt"

func main() {
	arr := []int{5, 2, 3, 1}        //1,2,3,5
	arr1 := []int{5, 1, 1, 2, 0, 0} //0,0,1,1,2,5
	fmt.Println(mergeSort(arr))
	fmt.Println(mergeSort(arr1))
}

//func sortArray(nums []int) []int {

//}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivot := len(nums) / 2
	left, right := mergeSort(nums[:pivot]), mergeSort(nums[pivot:])
	return merge(left, right)
}

func merge(l, r []int) []int {
	res := make([]int, 0, len(l)+len(r))
	fmt.Println(l, r, "начало")
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(res, r...)
		}
		if len(r) == 0 {
			return append(res, l...)
		}
		if l[0] < r[0] {
			res = append(res, l[0])
			l = l[1:]
		} else {
			res = append(res, r[0])
			r = r[1:]
		}
		fmt.Println(l, r)
	}
	return res
}
