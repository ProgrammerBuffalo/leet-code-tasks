package main

import "fmt"

var numLetters = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func main() {
	fmt.Println(letterCombinations("2376"))
}

func letterCombinations(digits string) []string {
	combinations := make([]string, 0, 4)
	backtracking(&combinations, digits, "")
	return combinations
}

func backtracking(combinations *[]string, digits, letters string) {
	if len(digits) == 1 {
		for _, r := range numLetters[rune(digits[0])] {
			*combinations = append(*combinations, letters+string(r))
		}
		return
	}
	for _, l := range numLetters[rune(digits[0])] {
		backtracking(combinations, digits[1:], letters+string(l))
	}
}
