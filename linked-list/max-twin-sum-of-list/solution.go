package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n1 := &ListNode{Val: 5}
	n2 := &ListNode{Val: 4}
	n3 := &ListNode{Val: 2}
	n4 := &ListNode{Val: 1}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	fmt.Println(pairSum(n1))
}

func pairSum(head *ListNode) int {
	left := head
	halfListSize := 1
	for right := head.Next.Next; ; right = right.Next.Next {
		if right == nil || right.Next == nil {
			break
		}
		left = left.Next
		halfListSize++
	}
	stack := make([]*ListNode, 0, halfListSize)
	for left = left.Next; left != nil; left = left.Next {
		stack = append(stack, left)
	}

	maxSum := 0
	left = head
	var rightTwin *ListNode
	for halfListSize > 0 {
		rightTwin = stack[halfListSize-1]
		stack = stack[:halfListSize-1]
		if rightTwin.Val+left.Val > maxSum {
			maxSum = rightTwin.Val + left.Val
		}
		halfListSize--
		left = left.Next
	}
	return maxSum
}
