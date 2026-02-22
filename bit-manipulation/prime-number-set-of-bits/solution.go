package main

import "fmt"

func main() {
	fmt.Println(countPrimeSetBits(6, 10))
}

func countPrimeSetBits(left int, right int) int {
	primes := 0
	for count, i := 0, left; i <= right; count, i = 0, i+1 {
		for j := i; j != 0; j = j >> 1 {
			if j&1 == 1 {
				count++
			}
		}
		if isPrime(count) {
			primes++
		}
	}
	return primes
}

func isPrime(n int) bool {
	switch n {
	case 2, 3, 5, 7, 11, 13, 17, 19:
		return true
	}
	return false
}
