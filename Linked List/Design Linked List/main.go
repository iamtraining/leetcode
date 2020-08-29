package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	circle := 0

	for {
		slow = slow.Next
		fast = fast.Next.Next
		circle++
		if slow == fast {
			break
		}
	}
	slow, fast = head, head
	for circle > 0 {
		fast = fast.Next
		circle--
	}
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
