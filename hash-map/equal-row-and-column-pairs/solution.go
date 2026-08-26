package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}))
}

func equalPairs(grid [][]int) int {
	rowM := make(map[string]int)

	var rowB strings.Builder
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			rowB.WriteString(strconv.Itoa(grid[i][j]))
			rowB.WriteByte(' ')
		}
		s := rowB.String()
		if count, ok := rowM[s]; ok {
			rowM[s] = count + 1
		} else {
			rowM[s] = 1
		}
		rowB.Reset()
	}

	eqCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			rowB.WriteString(strconv.Itoa(grid[j][i]))
			rowB.WriteByte(' ')
		}
		if count, ok := rowM[rowB.String()]; ok {
			eqCount += count
		}
		rowB.Reset()
	}
	return eqCount
}
