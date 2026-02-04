package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n1 := &ListNode{Val: 0}
	n2 := &ListNode{Val: 1}
	n3 := &ListNode{Val: 2}
	n4 := &ListNode{Val: 3}
	n5 := &ListNode{Val: 4}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n5.Next = n3

	fmt.Println(getIntersectionNode(n1, n5).Val)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	headACountToEnd, headBCountToEnd := 0, 0
	for itA := headA; itA != nil; itA = itA.Next {
		headACountToEnd++
	}
	for itB := headB; itB != nil; itB = itB.Next {
		headBCountToEnd++
	}

	for itA, itB := headA, headB; itA != nil && itB != nil; {
		if itA == itB {
			return itA
		} else if headACountToEnd > headBCountToEnd {
			itA = itA.Next
			headACountToEnd--
			continue
		} else if headACountToEnd < headBCountToEnd {
			itB = itB.Next
			headBCountToEnd--
			continue
		} else {
			itA = itA.Next
			itB = itB.Next
		}
	}

	return nil
}
