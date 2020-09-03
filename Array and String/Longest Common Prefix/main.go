package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix1([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix1([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix2([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix2([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	word := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], word) {
			word = word[:len(word)-1]
		}
	}
	return word
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || c != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
