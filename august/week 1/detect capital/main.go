package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("leetcode"))
	fmt.Println(detectCapitalUse("Google"))
	fmt.Println(detectCapitalUse("FlaG"))
}

func detectCapitalUse(word string) bool {
	letters := []rune(word)
	var lowercount, uppercount int
	for i := 0; i < len(letters); i++ {
		if unicode.IsUpper(letters[i]) == false {
			lowercount++
		} else {
			uppercount++
		}
	}
	if lowercount == len(letters) && uppercount == 0 {
		return true
	} else if unicode.IsUpper(letters[0]) == true && lowercount == len(letters)-1 {
		return true
	} else if uppercount == len(letters) && lowercount == 0 {
		return true
	}
	return false
}
