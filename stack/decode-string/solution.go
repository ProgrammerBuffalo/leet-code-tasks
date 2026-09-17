package main

import "fmt"

func main() {
	fmt.Println(decodeString("11[2[a]]"))
	fmt.Println(decodeString("2[2[y]pq]"))
	fmt.Println(decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef"))
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
}

func decodeString(s string) string {
	sSymbol := make([]byte, 0)
	sNum := make([]int, 0)

	for i := 0; i < len(s); i++ {
		if s[i] > 47 && s[i] < 58 {
			num := 0
			for i < len(s) {
				num *= 10
				if s[i] > 47 && s[i] < 58 {
					num += int(s[i] - '0')
					i++
				} else {
					num /= 10
					sNum = append(sNum, num)
					i--
					break
				}
			}
			continue
		}
		if s[i] == ']' {
			for j := len(sSymbol) - 1; j >= 0; j-- {
				if sSymbol[j] == '[' {
					repeated := repeatSymbols(sNum[len(sNum)-1], sSymbol[j+1:])
					sNum = sNum[:len(sNum)-1]
					sSymbol = sSymbol[:j]
					sSymbol = append(sSymbol, repeated...)
					break
				}
			}
		} else {
			sSymbol = append(sSymbol, s[i])
		}
	}
	return string(sSymbol)
}

func repeatSymbols(times int, symbolsToRepeat []byte) []byte {
	repeatedSymbols := make([]byte, 0, times*len(symbolsToRepeat))
	for i := 1; i <= times; i++ {
		for j := 0; j < len(symbolsToRepeat); j++ {
			repeatedSymbols = append(repeatedSymbols, symbolsToRepeat[j])
		}
	}
	return repeatedSymbols
}
