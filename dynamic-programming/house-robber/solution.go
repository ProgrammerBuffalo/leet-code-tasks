package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 1, 1, 2}))
}

func rob(nums []int) int {
	dp := make([]int, 3, len(nums)+3)

	for i := 0; i < len(nums); i++ {
		dp = append(dp, max(dp[i], dp[i+1])+nums[i])
	}

	return max(dp[len(dp)-1], dp[len(dp)-2])
}
