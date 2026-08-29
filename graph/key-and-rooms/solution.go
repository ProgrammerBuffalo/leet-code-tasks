package main

import "fmt"

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
}

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	stack := make([]int, 0)
	stack = append(stack, rooms[0]...)
	var curr int
	for len(stack) > 0 {
		curr = stack[len(stack)-1]
		if !visited[curr] {
			visited[curr] = true
			stack = append(stack, rooms[curr]...)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	for i := 1; i < len(visited); i++ {
		if !visited[i] {
			return false
		}
	}
	return true
}
