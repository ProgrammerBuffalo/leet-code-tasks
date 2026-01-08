package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := int((left + right) / 2)
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	if left >= len(nums)-1 {
		return left
	} else if right <= 0 {
		return left
	} else if nums[left] > target {
		return left
	} else {
		return right
	}
}
