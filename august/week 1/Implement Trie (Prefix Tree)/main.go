package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Insert("word")
	fmt.Println(obj.Search("word"))
	fmt.Println(obj.Search("world"))
	fmt.Println(obj.Startswith("word"))
	fmt.Println(obj.Startswith("world"))
	fmt.Println(obj.Startswith("ww"))
	fmt.Println(obj.Startswith("w"))
	fmt.Println(obj.Startswith("wo"))
	fmt.Println(obj.Startswith("wor"))
}

type Trie struct {
	next   [26]*Trie
	isWord bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curr := this
	for _, v := range word {
		i := v - 'a'
		if curr.next[i] == nil {
			curr.next[i] = &Trie{}
		}
		curr = curr.next[i]
	}
	curr.isWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, v := range word {
		i := v - 'a'
		if curr.next[i] == nil {
			return false
		}
		curr = curr.next[i]
	}
	return curr.isWord
}

func (this *Trie) Startswith(word string) bool {
	curr := this
	for _, v := range word {
		i := v - 'a'
		if curr.next[i] == nil {
			return false
		}
		curr = curr.next[i]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
