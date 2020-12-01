package main

import "fmt"

func main() {
	arr1 := []byte{'h', 'e', 'l', 'l', 'o'}
	arr2 := []byte{'h', 'a', 'n', 'n', 'a', 'H'}
	reverseString(arr1)
	reverseString(arr2)
	fmt.Println(string(arr1))
	fmt.Println(string(arr2))
}

func reverseString(s []byte) {
	helper(0, len(s)-1, s)
}

func helper(start, end int, s []byte) {
	if start >= end {
		return
	}
	tmp := s[end]
	s[end] = s[start]
	s[start] = tmp
	helper(start+1, end-1, s)
}
