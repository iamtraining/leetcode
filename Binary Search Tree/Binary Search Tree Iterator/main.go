package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var i BSTIterator
	i.push(root)
	return i
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	tmp := this.pop()
	this.push(tmp.Right)
	return tmp.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.stack) == 0 {
		return false
	}
	return true
}

func (i *BSTIterator) pop() *TreeNode {
	res := i.stack[len(i.stack)-1]
	i.stack = i.stack[:len(i.stack)-1]
	return res
}

func (i *BSTIterator) push(node *TreeNode) {
	for node != nil {
		i.stack = append(i.stack, node)
		node = node.Left
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
