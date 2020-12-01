package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return curr
}