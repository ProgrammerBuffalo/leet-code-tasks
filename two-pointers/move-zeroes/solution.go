package main

import "fmt"

func main() {
	nums := []int{1, 2}
	moveZeroes(nums)

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}

func moveZeroes(nums []int) {
	for idx, nonZeroIdx := 0, 0; nonZeroIdx < len(nums); nonZeroIdx++ {
		if nums[nonZeroIdx] == 0 {
			continue
		}
		temp := nums[idx]
		nums[idx] = nums[nonZeroIdx]
		nums[nonZeroIdx] = temp
		idx++
	}
}
