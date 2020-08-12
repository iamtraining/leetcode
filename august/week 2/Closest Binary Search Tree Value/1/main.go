package main

import "fmt"

func main() {
	var tree Node
	tree.insert([]int{4, 2, 5, 1, 3})
	fmt.Println(tree)
}

type Node struct {
	value       int
	left, right *Node
}

func (n *Node) insert(arr []int) {
	stack := reverse(arr)
	for len(stack) > 0 {
		i := len(stack) - 1
		if n == nil {
			n = &Node{value: stack[i]}
		} else if stack[i] > n.value {
			if n.right == nil {
				n.right = &Node{value: stack[i]}
			} else {
				n.right.insert(stack[i : i+1])
			}
		} else if stack[i] < n.value {
			if n.left == nil {
				n.left = &Node{value: stack[i]}
			} else {
				n.left.insert(stack[i : i+1])
			}
		}
		stack = stack[:i]
	}

}

func reverse(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		arr = append(arr, arr[i])
	}
	fmt.Println(arr[len(arr)/2:])
	return arr[len(arr)/2:]
}
