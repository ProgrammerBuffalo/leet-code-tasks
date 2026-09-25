package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: -3}

	root.Right = &TreeNode{Val: -2}

	fmt.Println(pathSum(root, -5))
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	targetSumCount := 0

	var innerDfs func(node *TreeNode, leftSum int)
	var outerDfs func(node *TreeNode)

	innerDfs = func(node *TreeNode, leftSum int) {
		if node == nil {
			return
		}

		if sum := leftSum - node.Val; sum == 0 {
			targetSumCount++
		}

		if node.Left != nil {
			innerDfs(node.Left, leftSum-node.Val)
		}
		if node.Right != nil {
			innerDfs(node.Right, leftSum-node.Val)
		}
	}

	outerDfs = func(node *TreeNode) {
		innerDfs(node, targetSum)
		if node.Left != nil {
			outerDfs(node.Left)
		}
		if node.Right != nil {
			outerDfs(node.Right)
		}
	}

	outerDfs(root)

	return targetSumCount
}
