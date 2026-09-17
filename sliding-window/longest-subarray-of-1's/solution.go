package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
	fmt.Println(longestSubarray([]int{1, 1, 0, 0, 0}))

	fmt.Println(longestSubarray([]int{1}))
	fmt.Println(longestSubarray([]int{0}))

	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))
	fmt.Println(longestSubarray([]int{1, 1, 1}))
}

func longestSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	left, right := 0, nums[0]
	latestZero := -1
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
		} else if latestZero != -1 {
			if maxLen < right-left {
				maxLen = right - left
			}
			left = latestZero + 1
			latestZero = i
		} else {
			latestZero = i
		}
		right = i

	}
	if maxLen < right-left {
		maxLen = right - left
	}
	return maxLen
}
