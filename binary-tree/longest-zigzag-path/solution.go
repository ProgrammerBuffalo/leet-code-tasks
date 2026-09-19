package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}

	root.Left = &TreeNode{Val: 2}

	root.Left.Left = &TreeNode{Val: 3}

	root.Left.Left.Right = &TreeNode{Val: 4}

	fmt.Println(longestZigZag(root))
}

func longestZigZag(root *TreeNode) int {
	maxDepth := 0
	var maxZigZag func(node *TreeNode, toLeft bool, depth int)
	maxZigZag = func(node *TreeNode, toLeft bool, depth int) {
		if toLeft {
			if node.Left != nil {
				maxZigZag(node.Left, !toLeft, depth+1)
			} else {
				if depth > maxDepth {
					maxDepth = depth
				}
			}
			if node.Right != nil {
				maxZigZag(node.Right, toLeft, 1)
			}
		} else {
			if node.Right != nil {
				maxZigZag(node.Right, !toLeft, depth+1)
			} else {
				if depth > maxDepth {
					maxDepth = depth
				}
			}
			if node.Left != nil {
				maxZigZag(node.Left, toLeft, 1)
			}
		}
	}

	maxZigZag(root, true, 0)

	return maxDepth
}
