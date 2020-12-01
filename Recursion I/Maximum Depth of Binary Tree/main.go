package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, count int) int {
	if root == nil {
		return count
	}
	count++
	left := helper(root.Left, count)
	if right := helper(root.Right, count); right > left {
		return right
	}
	return left
}
