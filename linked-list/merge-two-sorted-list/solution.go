package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	head2 := &ListNode{Val: 1, Next: &ListNode{Val: 4}}

	head3 := mergeTwoLists(head1, head2)

	for it := head3; it != nil; it = it.Next {
		fmt.Printf("%d ", it.Val)
	}
}

// Merge two list without returning new one will modify source lists
// This type of solution was decided for less memory use
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var list3 *ListNode

	if list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			list3 = list1
			list1 = list1.Next
		} else {
			list3 = list2
			list2 = list2.Next
		}
	}

	for it1, it2, it3 := list1, list2, list3; ; it3 = it3.Next {
		if it1 == nil {
			if it3 == nil {
				return it2
			}
			for ; it2 != nil; it2, it3 = it2.Next, it3.Next {
				it3.Next = it2
			}
			return list3
		} else if it2 == nil {
			if it3 == nil {
				return it1
			}
			for ; it1 != nil; it1, it3 = it1.Next, it3.Next {
				it3.Next = it1
			}
			return list3
		}

		if it1.Val < it2.Val {
			it3.Next = it1
			it1 = it1.Next
		} else {
			it3.Next = it2
			it2 = it2.Next
		}
	}

}
