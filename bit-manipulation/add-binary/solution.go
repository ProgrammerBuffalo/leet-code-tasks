package main

import "fmt"

func main() {
	a, b := "100", "110010"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	var res []rune

	for i, j, carry := len(a)-1, len(b)-1, 0; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		if i >= 0 {
			carry += int(a[i] - '0')
		}
		if j >= 0 {
			carry += int(b[j] - '0')
		}

		if carry%2 == 1 {
			res = append(res, '1')
			carry -= carry/2 + 1
		} else if carry != 0 {
			res = append(res, '0')
			carry--
		} else {
			res = append(res, '0')
		}

	}

	for i, j := len(res)-1, 0; i >= (len(res)-1)/2+1; i, j = i-1, j+1 {
		res[j], res[i] = res[i], res[j]
	}

	return string(res)
}
