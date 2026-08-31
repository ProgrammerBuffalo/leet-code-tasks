package main

import "fmt"

type Step struct {
	pos   []int
	count int
}

func main() {
	maze := [][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}
	fmt.Println(nearestExit(maze, []int{1, 2}))
	fmt.Println(maze)

	maze2 := [][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}}
	fmt.Println(nearestExit(maze2, []int{1, 0}))
	fmt.Println(maze2)

	maze4 := [][]byte{
		{'+', '.', '+', '+', '+', '+', '+'},
		{'+', '.', '+', '.', '.', '.', '+'},
		{'+', '.', '+', '.', '+', '.', '+'},
		{'+', '.', '.', '.', '+', '.', '+'},
		{'+', '+', '+', '+', '+', '.', '+'},
	}
	fmt.Println(nearestExit(maze4, []int{0, 1}))
	fmt.Println(maze4)

}

func nearestExit(maze [][]byte, entrance []int) int {
	queue := make([]Step, 0)
	queue = append(queue, Step{pos: entrance, count: 0})
	maze[entrance[0]][entrance[1]] = '+'
	var currentStep Step
	for len(queue) > 0 {
		currentStep = queue[0]
		queue = queue[1:]
		if currentStep.pos[0] > 0 && maze[currentStep.pos[0]-1][currentStep.pos[1]] == '.' {
			if currentStep.pos[0]-1 == 0 || currentStep.pos[1] == 0 || currentStep.pos[1] == len(maze[0])-1 {
				return currentStep.count + 1
			}
			maze[currentStep.pos[0]-1][currentStep.pos[1]] = '+'
			queue = append(queue, Step{pos: []int{currentStep.pos[0] - 1, currentStep.pos[1]}, count: currentStep.count + 1})
		}
		if currentStep.pos[0] < len(maze)-1 && maze[currentStep.pos[0]+1][currentStep.pos[1]] == '.' {
			if currentStep.pos[0]+1 == len(maze)-1 || currentStep.pos[1] == 0 || currentStep.pos[1] == len(maze[0])-1 {
				return currentStep.count + 1
			}
			maze[currentStep.pos[0]+1][currentStep.pos[1]] = '+'
			queue = append(queue, Step{pos: []int{currentStep.pos[0] + 1, currentStep.pos[1]}, count: currentStep.count + 1})
		}
		if currentStep.pos[1] > 0 && maze[currentStep.pos[0]][currentStep.pos[1]-1] == '.' {
			if currentStep.pos[1]-1 == 0 || currentStep.pos[0] == 0 || currentStep.pos[0] == len(maze)-1 {
				return currentStep.count + 1
			}
			maze[currentStep.pos[0]][currentStep.pos[1]-1] = '+'
			queue = append(queue, Step{pos: []int{currentStep.pos[0], currentStep.pos[1] - 1}, count: currentStep.count + 1})
		}
		if currentStep.pos[1] < len(maze[0])-1 && maze[currentStep.pos[0]][currentStep.pos[1]+1] == '.' {
			if currentStep.pos[1]+1 == len(maze[0])-1 || currentStep.pos[0] == 0 || currentStep.pos[0] == len(maze)-1 {
				return currentStep.count + 1
			}
			maze[currentStep.pos[0]][currentStep.pos[1]+1] = '+'
			queue = append(queue, Step{pos: []int{currentStep.pos[0], currentStep.pos[1] + 1}, count: currentStep.count + 1})
		}
	}
	return -1
}
