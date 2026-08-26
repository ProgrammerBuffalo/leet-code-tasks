package main

import "fmt"

func main() {
	fmt.Println(findDifference([]int{1, 2, 3}, []int{2, 4, 6}))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]struct{})
	for i := 0; i < len(nums1); i++ {
		set1[nums1[i]] = struct{}{}
	}
	set2 := make(map[int]struct{})
	for i := 0; i < len(nums2); i++ {
		set2[nums2[i]] = struct{}{}
	}
	ans := make([][]int, 2)
	for k1, _ := range set1 {
		if _, ok := set2[k1]; ok {
			delete(set1, k1)
			delete(set2, k1)
			continue
		}
		ans[0] = append(ans[0], k1)
	}
	for k2, _ := range set2 {
		ans[1] = append(ans[1], k2)
	}

	return ans
}
