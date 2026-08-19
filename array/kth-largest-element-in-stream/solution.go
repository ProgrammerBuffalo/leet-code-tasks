package main

import "fmt"

// I used a memory-efficient approach. However, for better time complexity, a MinHeap would be a better choice.
type KthLargest struct {
	nums []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	if len(nums) <= 0 {
		return KthLargest{nums: []int{-10000}, k: k}
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < k && len(nums)-1-j > i; j++ {
			if nums[i] >= nums[len(nums)-1-j] {
				nums[i], nums[len(nums)-1-j] = nums[len(nums)-1-j], nums[i]
			}
		}
	}

	if len(nums) < k {
		return KthLargest{
			nums: nums,
			k:    k,
		}
	}
	return KthLargest{
		nums: nums[len(nums)-k:],
		k:    k,
	}

}

func (this *KthLargest) Add(val int) int {
	if len(this.nums) < this.k {
		this.nums = append(this.nums, val)
		for i := len(this.nums) - 1; i > 0; i-- {
			if this.nums[i] < this.nums[i-1] {
				for j := i; j > 0; j-- {
					this.nums[j-1], this.nums[j] = this.nums[j], this.nums[j-1]
				}
			}
		}
		return this.nums[0]
	}

	for i := len(this.nums) - 1; i > 0; i-- {
		if this.nums[i] < val {
			for j := 0; j < i-1; j++ {
				this.nums[j] = this.nums[j+1]
			}
			this.nums[i-1], this.nums[i] = this.nums[i], val
			return this.nums[0]
		}
	}

	if this.nums[0] < val {
		this.nums[0] = val
	}

	return this.nums[0]
}

func main() {
	kth := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kth.Add(3))
	fmt.Println(kth.Add(5))
	fmt.Println(kth.Add(10))
	fmt.Println(kth.Add(9))
	fmt.Println(kth.Add(4))
}
