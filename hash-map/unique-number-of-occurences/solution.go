package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 1, 2, 2, 2, 3}))
}

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	set := make(map[int]struct{})

	for i := 0; i < len(arr); i++ {
		if count, ok := m[arr[i]]; !ok {
			m[arr[i]] = 1
		} else {
			m[arr[i]] = count + 1
		}
	}

	for _, val := range m {
		if _, ok := set[val]; ok {
			return false
		}
		set[val] = struct{}{}
	}
	return true
}
