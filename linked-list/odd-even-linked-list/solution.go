package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head2 := oddEvenList(
		&ListNode{Val: 1,
			Next: &ListNode{Val: 2}})
	for it := head2; it != nil; it = it.Next {
		fmt.Printf("%d ", it.Val)
	}
	fmt.Println()

	head1 := oddEvenList(
		&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}})
	for it := head1; it != nil; it = it.Next {
		fmt.Printf("%d ", it.Val)
	}
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	it := head

	headEven := it.Next
	tailEven := it.Next

	if it.Next == nil || it.Next.Next == nil {
		return head
	}

	var temp *ListNode
	var tailOdd *ListNode

	for it.Next != nil {
		temp = it.Next
		it.Next = it.Next.Next

		tailEven.Next = temp
		tailEven = tailEven.Next

		it = it.Next

		if it == nil {
			break
		}
		tailOdd = it
	}

	tailEven.Next = nil

	// tailOdd cannot be nil because of it.Next is not null
	tailOdd.Next = headEven

	return head
}
