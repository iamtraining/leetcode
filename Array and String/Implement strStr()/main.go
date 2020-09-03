package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStr1("hello", "ll"))
	fmt.Println(strStr1("aaaaaaa", "bba"))
	fmt.Println(strStr2("hello", "ll"))
	fmt.Println(strStr2("aaaaaaa", "bba"))
}

// Approach 1: Substring: Linear Time Slice

func strStr1(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if strings.Compare(haystack[i:i+len(needle)], needle) == 0 {
			return i
		}
	}
	return -1
}

// Approach 2: Two Pointers: Linear Time Slice

func strStr2(haystack string, needle string) int {
	ln, lhs, pn := len(needle), len(haystack), 0
	for pn < lhs-ln+1 {
		for pn < lhs-ln+1 && haystack[pn] != needle[0] {
			pn++
		}
		curr, pl := 0, 0
		for pl < ln && pn < lhs && haystack[pn] == needle[pl] {
			pn++
			pl++
			curr++
			if curr == ln {
				return pn - ln
			}
			pn = pn - curr + 1
		}
	}
	return -1
}
