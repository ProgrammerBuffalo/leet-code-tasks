package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	maxS := 0
	for s, left, right := 0, 0, len(height)-1; left != right; {
		if height[left] < height[right] {
			s = (right - left) * height[left]
			left++
		} else {
			s = (right - left) * height[right]
			right--
		}
		if s > maxS {
			maxS = s
		}

	}
	return maxS
}
