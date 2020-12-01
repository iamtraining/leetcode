package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	curr := head
	if curr != nil && curr.Next != nil {
		tmp := curr.Next.Val
		curr.Next.Val = curr.Val
		curr.Val = tmp
		swapPairs(curr.Next.Next)
	}

	return head
}
