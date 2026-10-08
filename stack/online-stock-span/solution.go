package main

import "fmt"

type StockSpanner struct {
	stack []int
}

func Constructor() StockSpanner {
	return StockSpanner{stack: make([]int, 0)}
}

func (s *StockSpanner) Next(price int) int {
	s.stack = append(s.stack, price)
	for i := len(s.stack) - 1; i >= 0; i-- {
		if s.stack[i] > price {
			return len(s.stack) - i - 1
		}
	}
	return len(s.stack)
}

func main() {
	stockSpanner := Constructor()
	fmt.Println(stockSpanner.Next(100)) // return 1
	fmt.Println(stockSpanner.Next(80))  // return 1
	fmt.Println(stockSpanner.Next(60))  // return 1
	fmt.Println(stockSpanner.Next(70))  // return 2
	fmt.Println(stockSpanner.Next(60))  // return 1
	fmt.Println(stockSpanner.Next(75))  // return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.
	fmt.Println(stockSpanner.Next(85))  // return 6
}
