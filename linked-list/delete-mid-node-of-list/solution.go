package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head1 := deleteMiddle(
		&ListNode{Val: 1})

	for it := head1; it != nil; it = it.Next {
		fmt.Printf("%d ", it.Val)
	}
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	left := head
	for right := head.Next.Next; ; right = right.Next.Next {
		if right == nil || right.Next == nil {
			break
		}
		left = left.Next
	}

	left.Next = left.Next.Next
	return head
}
