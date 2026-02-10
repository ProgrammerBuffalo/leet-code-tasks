package main

import (
	"fmt"
)

func main() {
	nums := []int{1, -1, 0, 2, -2, 1, -1, 0}

	triplets := threeSum(nums)

	for _, triplet := range triplets {
		fmt.Println(triplet)
	}
}

func threeSum(nums []int) [][]int {
	setNums := make(map[int]int)
	triplets := make(map[[3]int]struct{})

	for i := len(nums) - 1; i >= 0; i-- {
		setNums[nums[i]] = i
	}

	for x := 0; x < len(nums)-2; x++ {
		for y := x + 1; y < len(nums)-1; y++ {
			zNum := nums[x]*(-1) - nums[y]

			if z, ok := setNums[zNum]; ok {
				if z == x || z == y {
					continue
				}
				triplet := [3]int{nums[x], nums[y], zNum}
				sortTriplet(&triplet)
				triplets[triplet] = struct{}{}
			}
		}
	}

	res := make([][]int, 0, len(triplets))

	for triplet := range triplets {
		res = append(res, triplet[:])
	}

	return res
}

func sortTriplet(triplet *[3]int) {
	if triplet[0] > triplet[1] {
		triplet[0], triplet[1] = triplet[1], triplet[0]
	}
	if triplet[1] > triplet[2] {
		triplet[1], triplet[2] = triplet[2], triplet[1]
	}
	if triplet[0] > triplet[1] {
		triplet[0], triplet[1] = triplet[1], triplet[0]
	}
}
