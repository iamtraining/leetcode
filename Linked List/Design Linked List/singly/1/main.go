package main

import "fmt"

func main() {
	//obj := Constructor();
	//param_1 := obj.Get(index);
	//obj.AddAtHead(val);
	//obj.AddAtTail(val);
	//obj.AddAtIndex(index,val);
	//obj.DeleteAtIndex(index);
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	fmt.Println(obj.Get(0))
	fmt.Println(obj.Get(1)) // 2
	obj.DeleteAtIndex(1)
	fmt.Println(obj.Get(1)) // 3
}

type MyLinkedList struct {
	head   *Node
	length int
}

type Node struct {
	value int
	next  *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	count := 0
	for node := this.head; node != nil; node = node.next {
		if count == index {
			return node.value
		}
		count++
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.head = &Node{val, this.head}
	this.length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.head = &Node{value: val}

	}
	node := this.head
	for node.next != nil {
		node = node.next
	}
	node.next = &Node{value: val}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
	}
	count, prev := 0, this.head

	for node := this.head; node != nil; node = node.next {
		if count == index {
			prev.next = &Node{val, prev.next}
		}
		prev = node
		count++
	}
	if count == index {
		this.AddAtTail(val)
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}
	if index == 0 {
		if this.head.next != nil {
			this.head = &Node{this.head.next.value, this.head.next.next}
		}
		return
	}
	count, prev := 0, this.head

	for node := this.head; node != nil; node = node.next {
		if index == count {
			prev.next = node.next

		}
		prev = node
		count++
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
