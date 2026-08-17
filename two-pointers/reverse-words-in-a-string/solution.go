package main

import "fmt"

func main() {
	fmt.Println(reverseWords(" "))
	fmt.Println(reverseWords(" ab"))
	fmt.Println(reverseWords("   a  "))
	fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWords("a good   example"))
	fmt.Println(reverseWords("the sky is blue"))
}

func reverseWords(s string) string {
	reversed := make([]byte, 0, len(s))
	for left, right := len(s)-1, -1; left >= 0; left-- {
		if s[left] != ' ' {
			right = left
			for left >= 0 && s[left] != ' ' {
				left--
			}
			reversed = append(reversed, s[left+1:right+1]...)
			reversed = append(reversed, ' ')
		}
	}
	if len(reversed) == 0 {
		return ""
	}
	return string(reversed[:len(reversed)-1])
}
