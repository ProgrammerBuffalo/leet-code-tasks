package main

import "fmt"

func main() {
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, 0, len(cost)+1)

	dp = append(dp, cost...)
	dp = append(dp, 0)

	for i := 2; i < len(cost); i++ {
		if dp[i-1] >= dp[i-2] {
			dp[i] = cost[i] + dp[i-2]
			continue
		}
		dp[i] = cost[i] + dp[i-1]
	}

	return min(dp[len(dp)-3], dp[len(dp)-2])
}
