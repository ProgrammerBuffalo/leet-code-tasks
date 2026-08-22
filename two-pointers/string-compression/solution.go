package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch6 := []byte{'b', 'l', 'l', 'l', '4', '4', 'W', 'W', '&'}

	ch4 := []byte{'a', 'b', 'b', 'c'}
	ch5 := []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}

	ch3 := []byte{'a', 'b', 'c'}

	ch1 := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	ch2 := []byte{'a', 'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}

	fmt.Println(compress(ch6), string(ch6))

	fmt.Println(compress(ch4), string(ch4))
	fmt.Println(compress(ch5), string(ch5))

	fmt.Println(compress(ch1), string(ch1))
	fmt.Println(compress(ch2), string(ch2))
	fmt.Println(compress(ch3), string(ch3))

}

func compress(chars []byte) int {
	cursor, l, r := 0, 0, 0
	for {
		if chars[l] != chars[r] {
			cursor = addRepetition(chars, l, r, cursor)
			l = r
		} else if r == len(chars)-1 {
			if chars[l] == chars[r] {
				cursor = addRepetition(chars, l, r+1, cursor)
				break
			}
			break
		}
		if chars[l] == chars[r] {
			r++
		}
		if r >= len(chars) {
			if l+1 == r {
				chars[cursor] = chars[l]
				cursor++
			}
			break
		}
	}
	return cursor
}

func addRepetition(chars []byte, l, r, cursor int) int {
	chars[cursor] = chars[l]
	if r-l == 1 {
		return cursor + 1
	}
	repeatCount := strconv.Itoa(r - l)
	for i := 0; i < len(repeatCount); i++ {
		chars[cursor+1+i] = repeatCount[i]
	}
	cursor = cursor + 1 + len(repeatCount)
	return cursor
}
