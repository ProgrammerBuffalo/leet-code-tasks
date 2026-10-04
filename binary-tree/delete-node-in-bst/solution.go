package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 25}
	root.Left = &TreeNode{Val: 11}

	root.Right = &TreeNode{Val: 31}
	root.Right.Right = &TreeNode{Val: 32}
	root.Right.Left = &TreeNode{Val: 29}
	root.Right.Left.Right = &TreeNode{Val: 30}
	root.Right.Left.Left = &TreeNode{Val: 26}
	root.Right.Left.Left.Right = &TreeNode{Val: 27}
	root.Right.Left.Left.Right.Right = &TreeNode{Val: 28}

	inorder(root)
	fmt.Println("======")
	root = deleteNode(root, 25)
	inorder(root)

}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == key {
			return nil
		}
		return root
	}

	for cur, parent := root, root; cur != nil; {
		if cur.Val == key {
			if cur.Left == nil && cur.Right == nil {
				if parent.Left == cur {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
				return root
			}
			parent = cur
			prev := cur
			for cur.Left != nil || cur.Right != nil {
				if cur.Right != nil {
					prev, cur = changeByRightVal(parent, cur)
				} else {
					prev, cur = changeByLeftVal(parent, cur)
				}
				parent = cur
			}
			if prev.Left == cur {
				prev.Left = nil
			} else if prev.Right == cur {
				prev.Right = nil
			}
		}
		parent = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	return root
}

func changeByLeftVal(parent, node *TreeNode) (*TreeNode, *TreeNode) {
	prev := node
	node = node.Left
	for node.Right != nil {
		prev = node
		node = node.Right
	}
	parent.Val = node.Val
	return prev, node
}

func changeByRightVal(parent, node *TreeNode) (*TreeNode, *TreeNode) {
	prev := node
	node = node.Right
	for node.Left != nil {
		prev = node
		node = node.Left
	}
	parent.Val = node.Val
	return prev, node
}

func inorder(node *TreeNode) {
	if node == nil {
		return
	}

	inorder(node.Left)
	fmt.Println(node.Val)
	inorder(node.Right)
}
