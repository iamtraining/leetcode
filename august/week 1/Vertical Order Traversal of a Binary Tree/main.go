package main

func main() {
	root = new TreeNode(20,
                        new(TreeNode(7,
                                new(TreeNode(4, null, new(TreeNode(6)), new TreeNode(9))),
                        new TreeNode(35,
                                new TreeNode(31, new TreeNode(28), null),
                                new TreeNode(40, new(TreeNode(38)), new(TreeNode(52)))))

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
func verticalTraversal(root *TreeNode) [][]int {

}
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumDeep(tree TreeNode) int {
	var stack []int
	sum := 0

	stack = append(stack, tree.Val)

	for len(stack) > 0 {
		n := len(stack) - 1
		stack = stack[:n]
		sum += tree.Val
		if tree.Right != nil {
			stack = append(stack, tree.Right.Val)
		}
		if tree.Left != nil {
			stack = append(stack, tree.Left.Val)
		}
	}
	return sum
}
