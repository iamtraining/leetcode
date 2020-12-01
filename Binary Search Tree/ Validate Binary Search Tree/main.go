package main

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func bst(root *TreeNode) bool {
	return helper(root, nil, nil)
}

func helper(root, l, r *TreeNode) bool {
	if root == nil {
		return true
	}

	left, right := true, true
	if l != nil {
		left = l.Val <= root.Val
	}
	if r != nil {
		right = r.Val >= root.Val
	}

	return right && left && helper(root.Left, l, root) && helper(root.Right, root, r)
}
