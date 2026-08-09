package main

import "fmt"

func main() {
	fmt.Println(groupAnagrams([]string{"tea", "eat", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]byte][]string)
	count := [26]byte{0}
	for _, str := range strs {
		for _, ch := range str {
			count[ch-'a']++
		}
		m[count] = append(m[count], str)
		for i := 0; i < 26; i++ {
			count[i] = 0
		}
	}

	groups := make([][]string, 0, len(m))

	for _, anagrams := range m {
		groups = append(groups, anagrams)
	}

	return groups
}
