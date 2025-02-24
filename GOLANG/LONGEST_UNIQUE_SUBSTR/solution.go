package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	strLen := len(s)

	if strLen == 1 {
		return 1
	}

	// create set that will consist only ASCII - char's bytes
	set := make(map[byte]struct{})

	for i, left := 0, 0; i < strLen; i++ {
		// check is char's bytes already exists in set
		if _, ok := set[s[i]]; ok {
			if currLen := len(set); maxLen < currLen {
				maxLen = currLen
			}
			// delete leftmost item of set (SLIDING WINDOW alg.)
			delete(set, s[left])
			left++
			i--
			continue
		}
		// add unique char's byte
		set[s[i]] = struct{}{}
	}

	if currLen := len(set); maxLen < currLen {
		return currLen
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
