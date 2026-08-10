package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{4, 5, 1, 2, 3}))
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}))
	fmt.Println(findMin([]int{5, 1, 2, 3, 4}))
	fmt.Println(findMin([]int{4, 5, 6, 1, 2, 3}))
}

func findMin(nums []int) int {
	left, mid, right := 0, 0, len(nums)-1
	for {
		mid = (left + right) / 2

		if nums[mid] > nums[right] {
			if nums[left] > nums[right] {
				left = right
			}
			right = mid
		} else if nums[mid] < nums[left] {
			if nums[right] < nums[left] {
				right = left
			}
			left = mid
		} else {
			return nums[left]
		}

		if left-right == 1 || right-left == 1 {
			break
		}
	}
	return nums[left]
}
