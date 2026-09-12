package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: -100,
		Left: &TreeNode{
			Val: -200,
			Left: &TreeNode{
				Val: -20,
			},
			Right: &TreeNode{
				Val: -5,
			},
		},
		Right: &TreeNode{
			Val: -300,
			Left: &TreeNode{
				Val: -10,
			},
		},
	}

	fmt.Println(maxLevelSum(root))
}

func maxLevelSum(root *TreeNode) int {
	bfs := append(make([]*TreeNode, 0), root)
	maxSum, currentSum := root.Val, root.Val
	maxLvl := 1
	var size int
	var cursor *TreeNode
	for lvl := 1; len(bfs) > 0; lvl++ {
		size = len(bfs)
		for i := 0; i < size; i++ {
			cursor = bfs[i]
			currentSum += cursor.Val
			if cursor.Left != nil {
				bfs = append(bfs, cursor.Left)
			}
			if cursor.Right != nil {
				bfs = append(bfs, cursor.Right)
			}
		}
		if currentSum > maxSum {
			maxSum = currentSum
			maxLvl = lvl
		}
		currentSum = 0
		bfs = bfs[size:]
	}

	return maxLvl
}
