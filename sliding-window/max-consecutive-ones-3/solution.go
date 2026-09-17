package main

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))

	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}

func longestOnes(nums []int, k int) int {
	zeroIdxQueue := make([]int, 0)
	k = k + 1
	left, right := 0, 0
	longest := 0
	for i := 0; i < len(nums); i++ {
		right = i
		if nums[i] == 0 {
			zeroIdxQueue = append(zeroIdxQueue, i)
			k--
		}
		if k <= 0 {
			if longest < right-left {
				longest = right - left
			}
			left = zeroIdxQueue[0] + 1
			zeroIdxQueue = zeroIdxQueue[1:]
			k = 1
		}
	}
	if k >= 0 && longest < right-left+1 {
		longest = right - left + 1
	}
	return longest
}
