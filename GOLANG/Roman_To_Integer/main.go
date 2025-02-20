package main

import (
	"fmt"
)

func main() {
	sum := 0
	s := "MCMXCIV"

	m := map[rune]int{
		'I': 1, 'V': 5, 'X': 10,
		'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	runes := []rune(s)

	for i := 0; i < len(s); i++ {
		c1 := m[runes[i]]

		if i == len(s)-1 {
			sum += c1
			break
		}

		c2 := m[runes[i+1]]

		if c1 < c2 {
			sum += c2 - c1
			i++
			continue
		}

		sum += c1

	}
	fmt.Println(sum)
}
