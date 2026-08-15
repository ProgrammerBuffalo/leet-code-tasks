package main

import "fmt"

func main() {
	fmt.Println(gcdOfStrings("AAAAAAAAA", "AAACCC"))
}

func gcdOfStrings(str1 string, str2 string) string {
	j := 0
	g := gcd(len(str1), len(str2))
	if (str1 + str2) != (str2 + str1) {
		return ""
	}
	for i, d, nextDNth := 0, g, 0; ; i++ {
		if str1[i] != str2[j] {
			return ""
		}
		j = (j + 1) % d

		if i == len(str1)-1 {
			if j != 0 {
				for k, n := d-1, 0; k >= 1; k-- {
					if g%k == 0 {
						n++
					}
					if n == nextDNth {
						d = k
						nextDNth++
						j = 0
						i = 0
						break
					}
				}
			} else {
				return str2[:d]
			}
		}
	}

}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
