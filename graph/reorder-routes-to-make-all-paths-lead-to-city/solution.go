package main

import "fmt"

func main() {
	connections := [][]int{
		{0, 1},
		{1, 2},
		{1, 3},
	}
	fmt.Println(minReorder(4, connections))
}

func minReorder(n int, connections [][]int) int {
	list := make([][][2]int, n)
	/*
		{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}
		0 -> {0, 1}, {4, 0}
		1 -> {0, 1}, {1, 3}
		2 -> {2, 3}
		3 -> {1, 3}, {2, 3}
		4 -> {4, 0}, {4, 5}
		5 -> {4, 5}
	*/
	for i := 0; i < len(connections); i++ {
		list[connections[i][0]] = append(list[connections[i][0]], [2]int(connections[i]))
		list[connections[i][1]] = append(list[connections[i][1]], [2]int(connections[i]))
	}

	count := 0
	var bfs func(from int, fromTo [2]int)
	bfs = func(from int, curr [2]int) {
		if len(list[from]) == 1 {
			return
		}
		for j := 0; j < len(list[from]); j++ {
			if list[from][j][0] == curr[0] && list[from][j][1] == curr[1] {
				continue
			}
			if from == list[from][j][0] {
				count++
				bfs(list[from][j][1], list[from][j])
			} else {
				bfs(list[from][j][0], list[from][j])
			}
		}
	}

	for i := 0; i < len(list[0]); i++ {
		from := list[0][i][0]
		if from == 0 {
			count++
			from = list[0][i][1]
		}
		bfs(from, list[0][i])
	}

	return count
}
