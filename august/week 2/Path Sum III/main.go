package main

import "fmt"

func main() {
	var tree *TreeNode
	arr := []int{10, 5, -3, 3, 2, 11, 3, -2, 1}
	for i := 0; i < len(arr); i++ {
		tree = insert(tree, arr[i])
	}
	fmt.Println(pathSum(tree, 8))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insert(n *TreeNode, arg int) *TreeNode {
	if n == nil {
		n = &TreeNode{Val: arg}
	} else if arg < n.Val {
		n = insert(n.Left, arg)
	} else {
		n = insert(n.Right, arg)
	}
	return n
}

func pathSum(root *TreeNode, sum int) int {
	var arr []int
	count := 0
	helper(root, &count, sum, arr)
	return count
}

func helper(root *TreeNode, count *int, checkSum int, arr []int) {
	if root == nil {
		return
	}
	curr := 0
	arr = append(arr, root.Val)

	for i := len(arr) - 1; i >= 0; i-- {
		curr += arr[i]
		if curr == checkSum {
			*count++
		}
	}
	helper(root.Left, count, checkSum, arr)
	helper(root.Right, count, checkSum, arr)
}
