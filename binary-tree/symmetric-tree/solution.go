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
	fmt.Println(isSymmetric(root))
}

func isSymmetric(root *TreeNode) bool {
	return checkSymmetricNodes(root, root)
}

func checkSymmetricNodes(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return right == left
	}
	if left.Val != right.Val {
		return false
	}
	return checkSymmetricNodes(left.Left, right.Right) && checkSymmetricNodes(left.Right, right.Left)
}
