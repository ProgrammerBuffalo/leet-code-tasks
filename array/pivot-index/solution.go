package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{0, 0}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}

func pivotIndex(nums []int) int {
	prefixSum, suffixSum := 0, 0
	for i := len(nums) - 1; i > 0; i-- {
		suffixSum += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if prefixSum == suffixSum {
			return i
		}
		if i+1 < len(nums) {
			suffixSum -= nums[i+1]
		}
		prefixSum += nums[i]
	}
	return -1
}
