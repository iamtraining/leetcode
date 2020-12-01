package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	_, count := helper(root)
	return count
}

func helper(node *TreeNode) (bool, int) {
	count, ok := 0, true
	if node == nil {
		return true, 0
	}
	if node.Left == nil && node.Right == nil {
		return true, 1
	}

	if node.Left != nil {
		isLeft, leftCount := helper(node.Left)
		ok = (ok && isLeft) && (node.Val == node.Left.Val)
		count += leftCount
	}

	if node.Right != nil {
		isRight, rightCount := helper(node.Right)
		ok = (ok && isRight) && (node.Val == node.Right.Val)
		count += rightCount
	}

	if ok {
		count++
	}

	return ok, count

}
