package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 7)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	rotateTimes := k % len(nums)
	if rotateTimes == 0 {
		return
	}

	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	left, right = 0, rotateTimes-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	left, right = rotateTimes, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
