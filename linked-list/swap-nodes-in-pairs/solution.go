package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}

	list = swapPairs(list)

	for it := list; it != nil; it = it.Next {
		fmt.Printf("%d ", it.Val)
	}
}

func swapPairs(head *ListNode) *ListNode {
	var left, right *ListNode
	sentinel := &ListNode{Next: head}
	for it := sentinel; it != nil; it = it.Next {
		if it.Next != nil && it.Next.Next != nil {
			left = it.Next
			right = it.Next.Next.Next

			it.Next = it.Next.Next
			it.Next.Next = left
			left.Next = right

			it = it.Next
		}
	}

	return sentinel.Next
}
