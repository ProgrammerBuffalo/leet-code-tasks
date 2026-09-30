package main

import "fmt"

func main() {
	fmt.Println(increasingTriplet([]int{0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, -1, -1, -1, -1, 3}))
}

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	f, s := 0, 1
	for i := 0; i < len(nums); i++ {
		if nums[f] > nums[i] {
			temp := i
			for i = i + 1; i < len(nums)-1; i++ {
				if nums[f] < nums[s] && nums[s] < nums[i] {
					return true
				}
				if nums[i] > nums[temp] {
					s = i
					f = temp
					break
				} else if nums[i] < nums[temp] {
					temp = i
				}
			}
		}
		if nums[f] < nums[s] && nums[s] < nums[i] {
			return true
		}
		if nums[s] <= nums[f] || (nums[s] > nums[i] && nums[f] < nums[i]) {
			s = i
		}
	}

	return false
}
