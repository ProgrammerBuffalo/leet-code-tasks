package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	res := make([]int, len(spells))
	for i := 0; i < len(spells); i++ {
		l, r, mid := 0, len(potions)-1, 0
		for l < r {
			mid = (l + r) / 2
			if int64(potions[mid]*spells[i]) >= success {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		if r == len(potions)-1 {
			if int64(potions[r]*spells[i]) >= success {
				res[len(res)-1] = 1
			}
		} else {
			if int64(potions[l]*spells[i]) >= success {
				res[len(res)-1] = len(potions) - l
			} else {
				res[len(res)-1] = len(potions) - l - 1
			}
		}
	}
	return res
}
