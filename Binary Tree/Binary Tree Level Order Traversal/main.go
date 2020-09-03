package main

func main() {

}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	lvls, q, curr, lvl := [][]int{}, []*TreeNode{}, root, 0
	if root == nil {
		return lvls
	}
	q = append(q, curr)
	for len(q) != 0 {
		lvlen := len(q)
		for i := 0; i < lvlen; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			if len(lvls) <= lvl {
				lvls = append(lvls, []int{node.Val})
			} else {
				lvls[lvl] = append(lvls[lvl], node.Val)
			}
		}
		lvl++
	}
	return lvls
}
