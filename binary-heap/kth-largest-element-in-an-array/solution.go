package main

import "fmt"

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 20))
}

func findKthLargest(nums []int, k int) int {
	heap := make([]int, 0, k)
	for i := 0; i < len(nums); i++ {
		if len(heap) < k {
			heap = append(heap, nums[i])
			heapifyUp(heap)
		} else if nums[i] >= heap[0] {
			heap[0] = nums[i]
			heapifyDown(heap)
		}
	}

	return heap[0]
}

func heapifyUp(heap []int) {
	for j, n := (len(heap)-2)/2, len(heap)-1; ; n, j = j, (j-1)/2 {
		if heap[j] > heap[n] {
			heap[n], heap[j] = heap[j], heap[n]
		}
		if j <= 0 {
			break
		}
	}
}

func heapifyDown(heap []int) {
	for p, l, r := 0, 1, 2; l < len(heap); l, r = 2*p+1, 2*p+2 {
		if heap[p] >= heap[l] {
			if r < len(heap) {
				if heap[l] > heap[r] {
					heap[p], heap[r] = heap[r], heap[p]
					p = r
				} else {
					heap[p], heap[l] = heap[l], heap[p]
					p = l
				}
			} else {
				heap[p], heap[l] = heap[l], heap[p]
				break
			}
		} else if r < len(heap) && heap[p] >= heap[r] {
			heap[p], heap[r] = heap[r], heap[p]
			p = r
		} else {
			break
		}
	}
}
