package main

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		tmp := helper(root)
		root.Val = tmp.Val
		root.Right = deleteNode(root.Right, tmp.Val)
		return root
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}

func helper(root *TreeNode) *TreeNode {
	cur := root.Right
	for cur != nil && cur.Left != nil {
		cur = cur.Left
	}
	return cur
}
