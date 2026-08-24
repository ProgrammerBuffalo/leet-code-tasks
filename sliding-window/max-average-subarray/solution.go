package main

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}

func findMaxAverage(nums []int, k int) float64 {
	max := 0
	for i := 0; i < k; i++ {
		max += nums[i]
	}
	for sum, i := max, k; i < len(nums); i++ {
		sum -= nums[i-k]
		sum += nums[i]
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}
