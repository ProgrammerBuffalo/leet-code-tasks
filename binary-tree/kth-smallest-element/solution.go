package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 11}

	root.Left = &TreeNode{Val: 6}
	root.Right = &TreeNode{Val: 12}

	root.Left.Right = &TreeNode{Val: 9}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 10}

	fmt.Println(kthSmallest(root, 2))
}

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)

	// by using stack each node appends only one time for each iteration
	for it := root; ; {
		// this case occurs when the right node of the root is empty
		if it == nil {
			if len(stack) <= 0 {
				return -1
			}
			it = stack[len(stack)-1]
			k--
			if k == 0 {
				return it.Val
			}
			it = it.Right
			stack = stack[:len(stack)-1]
			continue
		}
		// if the left node is nil, then go instantly to the right node of root
		if it.Left != nil {
			stack = append(stack, it)
			it = it.Left
		} else {
			k--
			if k == 0 {
				return it.Val
			}
			it = it.Right
		}

	}
}
