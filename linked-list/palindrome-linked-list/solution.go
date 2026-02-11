package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 2,
							Next: &ListNode{Val: 1}}}}}}}

	fmt.Println(isPalindrome(head))
}

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	size := 0

	right := head
	for it := head; it != nil; it = it.Next {
		size++
	}

	for idx, mid := 0, int(size/2)+size%2; idx != mid; idx++ {
		right = right.Next
	}

	right = reverse(right)

	for left := head; ; left, right = left.Next, right.Next {
		if left.Val != right.Val {
			return false
		}
		if right == right.Next {
			break
		}
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	for cur, next := head, head; ; cur = next {
		if next.Next != nil {
			next = next.Next
			if next.Next != nil {
				next = next.Next
				cur.Next.Next = cur
				cur.Next, head = head, cur.Next
			} else {
				cur.Next.Next = cur
				cur.Next, head = head, cur.Next
				head = next
				break
			}
		} else {
			next.Next = head
			head = cur
			break
		}
	}

	return head
}
