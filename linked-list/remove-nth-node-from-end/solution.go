package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}

	head = removeNthFromEnd(head, 2)

	for it := head; it != nil; it = it.Next {
		fmt.Print(it.Val)
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	shift, prevOfRemove := head, &ListNode{Next: head}

	for i := 1; ; i++ {
		if i > n {
			prevOfRemove = prevOfRemove.Next
		}

		if shift.Next == nil {
			if head == prevOfRemove.Next {
				head = head.Next
			} else if prevOfRemove.Next != nil {
				prevOfRemove.Next = prevOfRemove.Next.Next
			}
			return head
		}

		shift = shift.Next
	}
}
