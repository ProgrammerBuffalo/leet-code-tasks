package main

import "fmt"

func main() {
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))
	fmt.Println(maxOperations([]int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}, 3))
}

func maxOperations(nums []int, k int) int {
	m := make(map[int]int)
	ops := 0
	for i := 0; i < len(nums); i++ {
		if count, ok := m[k-nums[i]]; ok {
			if count == 1 {
				delete(m, k-nums[i])
			} else {
				m[k-nums[i]]--
			}
			ops++
			continue
		}
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = 1
		} else {
			m[nums[i]]++
		}
	}
	return ops
}
