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

	fmt.Println(levelOrder(root))
}

func levelOrder(root *TreeNode) [][]int {
	levels := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	if root == nil {
		return levels
	}

	queue = append(queue, root)

	/* here I am using BFS via queue,
	   for example,
	   [ 6 | 12 | _ | _ | _ | _ | _ | _ ]
	   ^          ^
	   |          |
	   queue      nextQueue
	   (len = 2)  (len = 0)
	   (cap = 8)  (cap = 5)

	   when the cycle is finished, I just give to queue reference to the first element of nextQueue
	*/
	for len(queue) != 0 {
		nextQueue := queue[len(queue):]
		levelRow := make([]int, 0)
		for len(queue) != 0 {
			node := queue[0]
			queue = queue[1:]

			levelRow = append(levelRow, node.Val)

			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}

			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		levels = append(levels, levelRow)
		queue = nextQueue
	}

	return levels
}
