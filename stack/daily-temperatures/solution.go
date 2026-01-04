package main

import "fmt"

type Pair struct {
	temperature int
	index       int
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := make([]Pair, 0, len(temperatures))

	for i, t := range temperatures {
		if len(stack) > 0 && stack[len(stack)-1].temperature < t {
			last := stack[len(stack)-1]
			answer[last.index] = i - last.index
			stack = stack[:len(stack)-1]
			for len(stack) > 0 {
				if stack[len(stack)-1].temperature < t {
					last = stack[len(stack)-1]
					answer[last.index] = i - last.index
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
		}
		stack = append(stack, Pair{temperature: t, index: i})
	}

	return answer
}
