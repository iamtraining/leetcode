package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
}

func isPalindrome(s string) bool {
	sToLower := bytes.ToLower([]byte(s))
	var j int
	for _, ch := range sToLower {
		if ('a' <= ch && ch <= 'z') || ('0' <= ch && ch <= '9') {
			sToLower[j] = ch
			j++
		}
	}
	sToLower = sToLower[:j]
	for i, j := 0, len(sToLower)-1; i < j; i, j = i+1, j-1 {
		if sToLower[i] != sToLower[j] {
			return false
		}
	}
	return true
}
