package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil}}}

	reversedHead := reverseList(head)

	for it := reversedHead; it != nil; it = it.Next {
		fmt.Println(it.Val)
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	_, tail := reverseNode(head)
	return tail
}

func reverseNode(head *ListNode) (h *ListNode, t *ListNode) {
	if head.Next == nil {
		return head, head
	}

	next, tail := reverseNode(head.Next)

	next.Next = head
	head.Next = nil

	return head, tail
}
