package main

import "fmt"

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
}

func largestAltitude(gain []int) int {
	m := 0
	for sum, i := 0, 0; i < len(gain); i++ {
		sum += gain[i]
		if sum > m {
			m = sum
		}
	}
	return m
}
