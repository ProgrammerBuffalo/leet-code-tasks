package main

import "fmt"

func main() {
	fmt.Println(numIslands([][]byte{{'1', '1', '0'}, {'1', '0', '1'}, {'0', '0', '0'}}))

}

func numIslands(grid [][]byte) int {
	numOfIslands := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == '1' {
				dfs(x, y, grid)
				numOfIslands++
			}
		}
	}
	return numOfIslands
}

func dfs(x, y int, grid [][]byte) {
	if x >= len(grid) {
		return
	}
	if y >= len(grid[0]) {
		return
	}
	if grid[x][y] == '1' {
		grid[x][y] = '0'
		dfs(x+1, y, grid)
		dfs(x, y+1, grid)
		if x > 0 && grid[x-1][y] == '1' {
			dfs(x-1, y, grid)
		}
		if y > 0 && grid[x][y-1] == '1' {
			dfs(x, y-1, grid)
		}
	}

}
