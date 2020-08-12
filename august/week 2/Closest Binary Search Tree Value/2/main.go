package main

import (
	"fmt"
	"math"
)

func main() {
	var tree *Node
	arr := []int{4, 2, 5, 1, 3}
	for i := 0; i < len(arr); i++ {
		tree = insert(tree, arr[i])
	}
	min := traversal(tree, 3.714286)
	fmt.Println(min)
	fmt.Println(tree)
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func insert(n *Node, arg int) *Node {
	if n == nil {
		n = &Node{value: arg}
	} else if arg < n.value {
		n.left = insert(n.left, arg)
	} else {
		n.right = insert(n.right, arg)
	}

	return n
}

func traversal(n *Node, target float64) int {
	min := n.value
	for n != nil {
		if math.Abs(float64(n.value)-target) < math.Abs(float64(min)-target) {
			min = n.value
		}

		if target > float64(min) {
			n = n.left
		} else {
			n = n.right
		}
	}
	return min
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
