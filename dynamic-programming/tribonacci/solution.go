package main

import "fmt"

func main() {
	fmt.Println(tribonacci(5))
}

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	F := make([]int, 0, 38)
	F = append(F, 0)
	F = append(F, 1)
	F = append(F, 1)
	for i := 2; i < n; i++ {
		F = append(F, F[i-2]+F[i-1]+F[i])
	}
	return F[len(F)-1]
}
