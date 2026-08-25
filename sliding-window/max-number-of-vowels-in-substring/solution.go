package main

import "fmt"

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
}

func maxVowels(s string, k int) int {
	m := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	maxV := 0
	for i := 0; i < k; i++ {
		if _, ok := m[s[i]]; ok {
			maxV++
		}
	}
	for i, currMaxV := k, maxV; i < len(s); i++ {
		if _, ok := m[s[i-k]]; ok {
			currMaxV--
		}
		if _, ok := m[s[i]]; ok {
			currMaxV++
		}
		if currMaxV > maxV {
			maxV = currMaxV
		}
	}
	return maxV
}
