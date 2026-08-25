package main

import "fmt"

func main() {
	fmt.Println(removeStars("abc**"))
	fmt.Println(removeStars("leet**cod*e"))
}

func removeStars(s string) string {
	stack := make([]int32, 0, len(s))
	for _, ch := range s {
		if ch == 42 && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}
