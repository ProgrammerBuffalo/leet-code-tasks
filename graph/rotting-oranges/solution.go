package main

import "fmt"

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
}

func orangesRotting(grid [][]int) int {
	queue := make([][2]int, 0, len(grid)*len(grid[0]))
	freshCount := 0

	// accumulate all rotten fruits to queue
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}
	minsToRote := 0
	// BFS. Extract rotten fruit from the queue and check neighbor that can be rot
	for ; freshCount > 0; minsToRote++ {
		for i, itSize := 0, len(queue); i < itSize; i++ {
			rotten := queue[0]
			queue = queue[1:]
			if rotten[0] > 0 {
				if grid[rotten[0]-1][rotten[1]] == 1 {
					grid[rotten[0]-1][rotten[1]] = 2
					freshCount--
					queue = append(queue, [2]int{rotten[0] - 1, rotten[1]})
				}
			}
			if rotten[0]+1 < len(grid) {
				if grid[rotten[0]+1][rotten[1]] == 1 {
					grid[rotten[0]+1][rotten[1]] = 2
					freshCount--
					queue = append(queue, [2]int{rotten[0] + 1, rotten[1]})
				}
			}
			if rotten[1] > 0 {
				if grid[rotten[0]][rotten[1]-1] == 1 {
					grid[rotten[0]][rotten[1]-1] = 2
					freshCount--
					queue = append(queue, [2]int{rotten[0], rotten[1] - 1})
				}
			}
			if rotten[1]+1 < len(grid[0]) {
				if grid[rotten[0]][rotten[1]+1] == 1 {
					grid[rotten[0]][rotten[1]+1] = 2
					freshCount--
					queue = append(queue, [2]int{rotten[0], rotten[1] + 1})
				}
			}
		}
		if len(queue) == 0 && freshCount > 0 {
			return -1
		}
	}
	return minsToRote
}
