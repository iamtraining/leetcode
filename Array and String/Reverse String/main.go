package main

import "fmt"

func main() {
	arr := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(arr)
	fmt.Println(string(arr))
}

func reverseString(s []byte) {
	begin, end := 0, len(s)-1
	for begin < end {
		s[begin], s[end] = s[end], s[begin]
		begin++
		end--
	}
}
