package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	l3 := addTwoNumbers(l1, l2)

	for it := l3; it != nil; it = it.Next {
		fmt.Printf("%d ", it.Val)
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	res := &ListNode{}
	carry := 0
	itSum := 0

	for it1, it2, it3 := l1, l2, res; ; it3 = it3.Next {
		if it1 != nil {
			itSum += it1.Val
			it1 = it1.Next
		}
		if it2 != nil {
			itSum += it2.Val
			it2 = it2.Next
		}

		itSum += carry

		carry = itSum / 10
		it3.Val = itSum % 10
		if it1 != nil || it2 != nil {
			it3.Next = &ListNode{}
		} else if carry == 1 {
			it3.Next = &ListNode{Val: 1}
		} else {
			break
		}

		itSum = 0
	}

	return res
}
