package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 {
			return n-1 <= 0
		}
		return n <= 0
	}
	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		n--
		flowerbed[0] = 1
	}
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			n--
			flowerbed[i] = 1
		}
	}
	if flowerbed[len(flowerbed)-2] == 0 && flowerbed[len(flowerbed)-1] == 0 {
		n--
		flowerbed[0] = 1
	}

	return n <= 0
}
