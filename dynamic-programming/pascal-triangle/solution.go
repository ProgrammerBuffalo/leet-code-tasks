package main

import "fmt"

func main() {
	fmt.Println(generate(6))
}

func generate(numRows int) [][]int {
	triangle := make([][]int, 0, numRows)
	triangle = append(triangle, []int{1})

	if numRows <= 1 {
		return triangle
	}

	for i := 2; i <= numRows; i++ {
		row := make([]int, 0, i)
		row = append(row, 1)
		for j := 2; j < i; j++ {
			row = append(row, triangle[i-2][j-2]+triangle[i-2][j-1])
		}
		row = append(row, 1)
		triangle = append(triangle, row)
	}

	return triangle
}

// 1           = 1
// 1 1         = 2
// 1 2 1       = 3
// 1 3 3 1     = 4
// 1 4 6 4 1   = 5
