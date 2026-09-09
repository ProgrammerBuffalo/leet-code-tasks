package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 9))
}

func combinationSum3(k int, n int) [][]int {
	results := make([][]int, 0)
	result := make([]int, 0, k)

	var findCombinations func(from int, leftK int, leftN int)

	findCombinations = func(current int, leftK int, leftN int) {
		for i := current; i <= 9; i++ {
			if leftN-i < 0 {
				return
			}
			if leftK == 1 && leftN-i == 0 {
				result = append(result, i)
				results = append(results, result)

				result = make([]int, 0)
				result = append(result, results[len(results)-1][:k-1]...)
				return
			}
			if leftN-i > i {
				result = append(result, i)
				findCombinations(i+1, leftK-1, leftN-i)
				result = result[:len(result)-1]
			}
		}
	}

	findCombinations(1, k, n)
	return results
}
