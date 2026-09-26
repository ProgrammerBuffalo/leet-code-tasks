package main

import "fmt"

func main() {
	fmt.Println(findCircleNum([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}))
}

// 		1  0  0  1
// 		0  1  1  0
// 		0  1  1  1
// 		1  0  1  1

func findCircleNum(isConnected [][]int) int {
	circleNum := 0
	dfs := make([]int, 0)
	for i := 0; i < len(isConnected); i++ {
		if isConnected[i][i] == 0 {
			continue
		}
		isConnected[i][i] = 0
		for j := i + 1; j < len(isConnected); j++ {
			if isConnected[i][j] == 1 {
				dfs = append(dfs, j)
			}
		}
		for len(dfs) > 0 {
			curr := dfs[len(dfs)-1]
			dfs = dfs[:len(dfs)-1]
			isConnected[curr][curr] = 0
			for j := 0; j < len(isConnected); j++ {
				if isConnected[curr][j] == 1 && isConnected[j][j] != 0 {
					dfs = append(dfs, j)
				}
			}
		}
		circleNum++
	}
	return circleNum
}
