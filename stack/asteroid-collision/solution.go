package main

import "fmt"

func main() {
	fmt.Println(asteroidCollision([]int{-2, -2, 1, -2}))
	fmt.Println(asteroidCollision([]int{3, 5, -6, 2, -1, 4}))
}

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))
	stack = append(stack, asteroids[0])
	for i := 1; i < len(asteroids); i++ {
		if len(stack) > 0 && stack[len(stack)-1] > 0 && asteroids[i] < 0 {
			diff := 0
			for {
				diff = stack[len(stack)-1] + asteroids[i]
				if diff >= 0 {
					if diff == 0 {
						stack = stack[:len(stack)-1]
					}
					break
				}
				stack = stack[:len(stack)-1]
				if len(stack) == 0 || stack[len(stack)-1] < 0 && asteroids[i] < 0 {
					stack = append(stack, asteroids[i])
					break
				}
			}
		} else {
			stack = append(stack, asteroids[i])
		}
	}
	return stack
}
