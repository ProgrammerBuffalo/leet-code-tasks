package main

import "fmt"

func main() {
	fmt.Println(findPeakElementWith_O_logn([]int{1, 0}))
}

func findPeakElementWith_O_logn(nums []int) int {
	l, r := 0, len(nums)-1
	var mid int
	for l < r {
		mid = (l + r) / 2
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func findPeakElementWith_O_n(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 0 {
		return -1
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}

	l, r := len(nums)/2-1, len(nums)/2+1
	if nums[l] <= nums[l+1] && nums[r] <= nums[r-1] {
		return l + 1
	}
	for l != 0 || r != len(nums)-1 {
		if l == 0 {
			continue
		} else if nums[l] < nums[l-1] {
			l = l - 1
		} else if nums[l+1] <= nums[l] {
			return l
		}
		if r == len(nums)-1 {
			continue
		} else if nums[r] < nums[r+1] {
			r = r + 1
		} else if nums[r-1] <= nums[r] {
			return r
		}
	}
	return -1
}
