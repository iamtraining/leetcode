package main

func main() {

}

type MyHashSet struct {
	Map map[int]bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{Map: make(map[int]bool)}
}

func (this *MyHashSet) Add(key int) {
	if _, ok := this.Map[key]; !ok {
		this.Map[key] = true
	}
}

// Remove ...
func (this *MyHashSet) Remove(key int) {
	if _, ok := this.Map[key]; ok {
		delete(this.Map, key)
	}
}

/** Returns true if this set contains the specified element */
// Contains...
func (this *MyHashSet) Contains(key int) bool {
	if _, ok := this.Map[key]; ok {
		return true
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
