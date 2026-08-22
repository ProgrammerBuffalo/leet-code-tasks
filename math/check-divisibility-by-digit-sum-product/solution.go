package main

import "fmt"

func main() {
	fmt.Println(checkDivisibility(10))
	fmt.Println(checkDivisibility(99))
	fmt.Println(checkDivisibility(23))
}

func checkDivisibility(n int) bool {
	sum, product := 0, 1
	digit := 0
	for i := 1; i <= n; i *= 10 {
		digit = (n % (i * 10)) / i
		sum += digit
		product *= digit
	}
	return n%(sum+product) == 0
}
