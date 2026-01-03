package main

import "fmt"

func main() {
	fmt.Println(isValid("{{}}()[]"))
	fmt.Println(isValid("{{}]()[]"))
	fmt.Println(isValid("{}()[]"))
	fmt.Println(isValid("((((()))))"))
	fmt.Println(isValid("(("))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0, len(s)/2)

	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
		} else if r == ')' || r == ']' || r == '}' {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			switch last {
			case '(':
				if r != ')' {
					return false
				}
			case '[':
				if r != ']' {
					return false
				}
			case '{':
				if r != '}' {
					return false
				}
			}
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
