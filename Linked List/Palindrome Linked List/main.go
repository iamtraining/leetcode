package main

func main() {

}

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) {

}

func isPalindrome1(head *ListNode) bool {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

func isPalindrome3(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//isOdd
	if fast != nil {
		slow = slow.Next
	}
	slow = reverse(slow)
	for slow != nil {
		if head.Val == slow.Val {
			head, slow = head.Next, slow.Next
			continue
		}
		return false
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}
