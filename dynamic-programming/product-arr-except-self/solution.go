package main

import "fmt"

func main() {
	nums := make([]int, 0)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)

	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	dp := make([]int, 0, len(nums)+2)

	dp = append(dp, 1)
	for i := 1; i < len(nums); i++ {
		dp = append(dp, dp[i-1]*nums[i-1])
	}

	nums = append(nums, 1)
	for i := len(nums) - 2; i >= 0; i-- {
		nums[i] = nums[i] * nums[i+1]
	}

	for i := 0; i < len(dp); i++ {
		nums[i] = nums[i+1] * dp[i]
	}

	return nums[:len(nums)-1]
}
