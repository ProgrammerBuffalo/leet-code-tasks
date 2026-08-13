package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{4, 2, 8}, 8))
}

func combinationSum(candidates []int, target int) [][]int {
	results := make([][]int, 0)
	currentResult := make([]int, 0)

	var backtrack func(idx, leftSum int)

	backtrack = func(idx, leftSum int) {
		if idx >= len(candidates) {
			return
		}

		if leftSum == 0 {
			results = append(results, append([]int(nil), currentResult...))
			return
		}

		if leftSum-candidates[idx] >= 0 {
			currentResult = append(currentResult, candidates[idx])
			backtrack(idx, leftSum-candidates[idx])
			currentResult = currentResult[:len(currentResult)-1]
		}

		backtrack(idx+1, leftSum)

	}

	backtrack(0, target)
	return results
}
