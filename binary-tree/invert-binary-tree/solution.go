package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	invertTree(root)
	printTree(root)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	invertSide(root.Left)
	invertSide(root.Right)

	root.Left, root.Right = root.Right, root.Left

	return root
}

func invertSide(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		invertTree(root.Left)
	}

	if root.Right != nil {
		invertTree(root.Right)
	}

	root.Left, root.Right = root.Right, root.Left
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)

	printTree(root.Left)
	printTree(root.Right)
}
