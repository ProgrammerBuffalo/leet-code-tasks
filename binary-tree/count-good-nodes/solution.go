package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type StackItem struct {
	It  *TreeNode
	Max int
}

func main() {
	root := &TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: -4,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
			},
		},
		Right: &TreeNode{
			Val: -2,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: -2,
					Left: &TreeNode{
						Val: -1,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: -3,
					},
				},
			},
			Right: &TreeNode{
				Val: -2,
				Right: &TreeNode{
					Val: -2,
					Left: &TreeNode{
						Val: -4,
					},
					Right: &TreeNode{
						Val: -3,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: -3,
						},
					},
				},
			},
		},
	}
	fmt.Println(goodNodes(root))
}

func goodNodes(root *TreeNode) int {
	stack := make([]StackItem, 0)
	maxX := root.Val
	count := 0
	var item StackItem
	for it := root; it != nil; {
		if it.Val >= maxX {
			maxX = it.Val
			count++
		}
		if it.Left != nil {
			stack = append(stack, StackItem{It: it, Max: maxX})
			it = it.Left
			continue
		} else if it.Right != nil {
			it = it.Right
			continue
		}
		if len(stack) == 0 {
			break
		}
		for item = stack[len(stack)-1]; ; item = stack[len(stack)-1] {
			maxX = item.Max
			stack = stack[:len(stack)-1]
			it = item.It.Right
			if len(stack) == 0 || it != nil {
				break
			}
		}

	}
	return count
}
