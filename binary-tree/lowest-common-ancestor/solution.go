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

	fmt.Println(lowestCommonAncestor(root, root.Left.Right.Left, root.Left).Val)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pSet := make(map[int]*TreeNode)
	var dfsP func(root *TreeNode, descendant *TreeNode) bool
	dfsP = func(root *TreeNode, descendant *TreeNode) bool {
		if root == descendant {
			pSet[root.Val] = root
			return true
		}
		if root == nil {
			return false
		}
		if root.Left != nil {
			if dfsP(root.Left, descendant) {
				pSet[root.Val] = root
				return true
			}
		}
		if root.Right != nil {
			if dfsP(root.Right, descendant) {
				pSet[root.Val] = root
				return true
			}
		}
		return false
	}

	dfsP(root, p)
	var lowestAncestor *TreeNode
	var dfsQ func(root *TreeNode, descendant *TreeNode) bool
	dfsQ = func(root *TreeNode, descendant *TreeNode) bool {
		if root == descendant {
			if _, ok := pSet[root.Val]; ok {
				lowestAncestor = pSet[root.Val]
				return false
			}
			return true
		}
		if root == nil {
			return false
		}
		if root.Left != nil {
			if dfsQ(root.Left, descendant) {
				if _, ok := pSet[root.Val]; ok {
					lowestAncestor = pSet[root.Val]
					return false
				}
				return true
			}
		}
		if root.Right != nil {
			if dfsQ(root.Right, descendant) {
				if _, ok := pSet[root.Val]; ok {
					lowestAncestor = pSet[root.Val]
					return false
				}
				return true
			}
		}
		return false
	}

	dfsQ(root, q)

	return lowestAncestor
}
