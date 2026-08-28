package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(leafSimilar(root1, root2))
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	stack1, stack2 := make([]*TreeNode, 0), make([]*TreeNode, 0)
	it1, it2 := root1, root2
	for {
		if it1 == nil && it2 == nil {
			return true
		} else if it1 == nil || it2 == nil {
			return false
		}
		if it1.Left != nil {
			stack1 = append(stack1, it1)
			it1 = it1.Left
		} else if it1.Right != nil {
			it1 = it1.Right
		} else {
			for {
				if it2.Left != nil {
					stack2 = append(stack2, it2)
					it2 = it2.Left
				} else if it2.Right != nil {
					it2 = it2.Right
				} else {
					if it1.Val != it2.Val {
						return false
					}
					for len(stack2) > 0 {
						it2 = stack2[len(stack2)-1]
						stack2 = stack2[:len(stack2)-1]
						if it2.Right != nil {
							break
						}
					}
					for len(stack1) > 0 {
						it1 = stack1[len(stack1)-1]
						stack1 = stack1[:len(stack1)-1]
						if it1.Right != nil {
							break
						}
					}
					it1 = it1.Right
					it2 = it2.Right

					break

				}
			}
		}
	}
}
