package main

import "fmt"

func main() {
	obj := Constructor()
	obj.AddWord("word")
	fmt.Println(obj.Search("word"))
	fmt.Println(obj.Search("wo.d"))
	fmt.Println(obj.Search("...d"))
	fmt.Println(obj.Search(".o.d"))
	fmt.Println(obj.Search(".o.d"))
	fmt.Println(obj.Search("z"))
	fmt.Println(obj.Search("d"))
}

type WordDictionary struct {
	next   [26]*WordDictionary
	isWord bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	curr := this
	for _, v := range word {
		i := v - 'a'
		if curr.next[i] == nil {
			curr.next[i] = &WordDictionary{}
		}
		curr = curr.next[i]
	}
	curr.isWord = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	curr := this
	for i := range word {
		if word[i] == '.' {
			for _, ch := range curr.next {
				if ch != nil && ch.Search(word[i+1:]) {
					return true
				}
			}
			return false
		}
		ind := word[i] - 'a'
		if curr.next[ind] == nil {
			return false
		}
		curr = curr.next[ind]
	}
	return curr.isWord
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
