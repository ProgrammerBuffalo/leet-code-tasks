package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n1 := &ListNode{}

	n1.Next = &ListNode{Next: &ListNode{Next: &ListNode{Next: &ListNode{Next: n1}}}}

	fmt.Println(hasCycle(n1))
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	for slow, fast := head, head.Next; ; {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			return true
		}
	}
}
