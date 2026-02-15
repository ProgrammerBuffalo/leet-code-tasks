package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 8}

	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 10}

	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 6}

	root.Left.Right.Left = &TreeNode{Val: 4}
	root.Left.Right.Right = &TreeNode{Val: 7}

	root.Right.Right = &TreeNode{Val: 14}
	root.Right.Right.Left = &TreeNode{Val: 13}

	fmt.Print(inorderTraversal(root))
}

func inorderTraversal(root *TreeNode) []int {
	var arr []int
	var stack []*TreeNode

	for it := root; it != nil || len(stack) > 0; {
		if it == nil && len(stack) > 0 {
			arr = append(arr, stack[len(stack)-1].Val)
			it = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
			continue
		}

		if it.Left != nil {
			stack = append(stack, it)
			it = it.Left
		} else {
			arr = append(arr, it.Val)
			it = it.Right
		}
	}
	return arr
}
