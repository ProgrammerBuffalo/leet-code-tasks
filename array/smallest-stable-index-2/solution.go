package main

import "fmt"

func main() {
	fmt.Println(firstStableIndex([]int{5, 0, 1, 4}, 3))
	fmt.Println(firstStableIndex([]int{3, 2, 1}, 1))
	fmt.Println(firstStableIndex([]int{0}, 0))
}

func firstStableIndex(nums []int, k int) int {
	suffixMin := make([]int, len(nums))
	prefixMax := nums[0]
	for i, minimum := len(nums)-1, nums[len(nums)-1]; i >= 0; i-- {
		if nums[i] <= minimum {
			minimum = nums[i]
		}
		suffixMin[i] = minimum
	}
	for i := 0; i < len(nums); i++ {
		if prefixMax < nums[i] {
			prefixMax = nums[i]
		}
		if prefixMax-suffixMin[i] <= k {
			return i
		}
	}
	return -1
}
