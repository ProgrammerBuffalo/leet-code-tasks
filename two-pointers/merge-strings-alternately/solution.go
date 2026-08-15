package main

import "fmt"

func main() {
	fmt.Println(mergeAlternately("ab", "pqrs"))
}

func mergeAlternately(word1 string, word2 string) string {
	mergedStr := make([]byte, 0)

	for i, j := 0, 0; len(word1) > i || len(word2) > j; {
		if len(word1) > i {
			mergedStr = append(mergedStr, word1[i])
			i++
		}
		if len(word2) > j {
			mergedStr = append(mergedStr, word2[j])
			j++
		}
	}

	return string(mergedStr)
}
