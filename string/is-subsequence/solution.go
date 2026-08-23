package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "afdbsdc"))
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}

	sIdx := 0
	for tIdx := 0; tIdx < len(t); tIdx++ {
		if s[sIdx] == t[tIdx] {
			sIdx++
		}
		if sIdx == len(s) {
			return true
		}
	}
	return false
}
