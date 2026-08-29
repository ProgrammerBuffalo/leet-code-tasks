package main

import (
	"fmt"
	"math"
)

var pick = 777

func main() {
	fmt.Println(guessNumber(math.MaxInt32))
}

func guessNumber(n int) int {
	left, mid, right := 0, n/2, n
	guessRes := 0
	for {
		guessRes = guess(mid)
		if guessRes == -1 {
			right = mid - 1
		} else if guessRes == 1 {
			left = mid + 1
		} else {
			return mid
		}
		mid = (left + right) / 2
	}
}

func guess(guessNum int) int {
	if guessNum > pick {
		return -1
	}
	if guessNum < pick {
		return 1
	}
	return 0
}
