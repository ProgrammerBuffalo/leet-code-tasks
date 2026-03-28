package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(maxDepth(root))
}

func maxDepth(root *TreeNode) int {
	return maxDepthFromRoot(root, 1)
}

func maxDepthFromRoot(root *TreeNode, depth int) int {
	if root == nil {
		return depth - 1
	}

	leftDepth := maxDepthFromRoot(root.Left, depth+1)
	rightDepth := maxDepthFromRoot(root.Right, depth+1)

	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}
