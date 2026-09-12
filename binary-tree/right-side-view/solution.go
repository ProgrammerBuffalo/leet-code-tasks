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
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(rightSideView(root))
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	bfs := make([]*TreeNode, 0)

	bfs = append(bfs, root)
	var size int
	var curr *TreeNode
	for len(bfs) > 0 {
		res = append(res, bfs[len(bfs)-1].Val)
		size = len(bfs)
		for i := 0; i < size; i++ {
			curr = bfs[i]
			if curr.Left != nil {
				bfs = append(bfs, curr.Left)
			}
			if curr.Right != nil {
				bfs = append(bfs, curr.Right)
			}
		}
		bfs = bfs[size:]
	}

	return res
}
