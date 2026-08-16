package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	results := make([]bool, 0, len(candies))
	for i := 0; i < len(candies); i++ {
		if maxCandies < candies[i] {
			maxCandies = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		results = append(results, (candies[i]+extraCandies) >= maxCandies)
	}
	return results
}
