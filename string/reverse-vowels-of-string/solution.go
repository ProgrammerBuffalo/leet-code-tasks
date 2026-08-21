package main

import "fmt"

var vowels = map[byte]struct{}{
	'a': {},
	'A': {},
	'e': {},
	'E': {},
	'i': {},
	'I': {},
	'o': {},
	'O': {},
	'u': {},
	'U': {},
}

func reverseVowels(s string) string {
	str := []byte(s)
	for i, j := 0, len(s)-1; i < j; {
		if _, ok := vowels[s[i]]; ok {
			for ; i != j; j-- {
				if _, ok = vowels[s[j]]; ok {
					str[i], str[j] = str[j], str[i]
					break
				}
			}
		} else if _, ok = vowels[s[j]]; ok {
			for ; i != j; i++ {
				if _, ok = vowels[s[i]]; ok {
					str[i], str[j] = str[j], str[i]
					break
				}
			}
		}
		i++
		j--
	}
	return string(str)
}

func main() {
	fmt.Println(reverseVowels("IceCreAm"))
}
