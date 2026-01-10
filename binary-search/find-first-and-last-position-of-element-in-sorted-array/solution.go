package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 9}, 8))
}

func searchRange(nums []int, target int) []int {
	indexes := []int{-1, -1}
	mid := findMid(0, len(nums)-1, nums, target)

	if mid == -1 {
		return indexes
	}

	for left := mid; ; {
		if left > 0 && nums[left-1] == nums[left] {
			left = findMid(0, left-1, nums, target)
		} else {
			indexes[0] = left
			break
		}
	}

	for right := mid; ; {
		if right < len(nums)-1 && nums[right+1] == nums[right] {
			right = findMid(right+1, len(nums)-1, nums, target)
		} else {
			indexes[1] = right
			break
		}
	}

	return indexes
}

func findMid(left int, right int, nums []int, target int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
