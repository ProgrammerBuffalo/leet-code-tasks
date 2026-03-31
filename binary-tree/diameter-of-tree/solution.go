package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(diameterOfBinaryTree(root))
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := maxDepthOfSide(root.Left)
	rDepth := maxDepthOfSide(root.Right)

	diameter := lDepth + rDepth
	nextDiameter := 0
	if lDepth > rDepth {
		nextDiameter = diameterOfBinaryTree(root.Left)
	} else {
		nextDiameter = diameterOfBinaryTree(root.Right)
	}

	if diameter > nextDiameter {
		return diameter
	}
	return nextDiameter
}

func maxDepthOfSide(root *TreeNode) int {
	if root == nil {
		return 0
	}
	rDepth := maxDepthOfSide(root.Right)
	lDepth := maxDepthOfSide(root.Left)

	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}
