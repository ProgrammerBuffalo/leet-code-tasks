package main

import (
	"fmt"
	"sync"
)

type ZeroEvenOdd struct {
	n int

	evenCh chan int
	oddCh  chan int
	zeroCh chan int
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeroCh := make(chan int, 1)
	zeroCh <- 1
	return &ZeroEvenOdd{
		n: n,

		evenCh: make(chan int, 1),
		oddCh:  make(chan int, 1),
		zeroCh: zeroCh,
	}
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for v := range z.zeroCh {
		printNumber(0)
		if v%2 == 0 {
			z.evenCh <- v
		} else {
			z.oddCh <- v
		}
	}
	close(z.oddCh)
	close(z.evenCh)
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for v := range z.evenCh {
		printNumber(v)
		if v == z.n {
			close(z.zeroCh)
			return
		}
		v++

		z.zeroCh <- v

	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for v := range z.oddCh {
		printNumber(v)
		if v == z.n {
			close(z.zeroCh)
			return
		}
		v++

		z.zeroCh <- v
	}
}

func main() {
	wg := &sync.WaitGroup{}
	z := NewZeroEvenOdd(10)
	wg.Add(3)

	go func(z *ZeroEvenOdd) {
		defer wg.Done()
		z.Zero(func(z int) { fmt.Println(z) })
	}(z)
	go func(z *ZeroEvenOdd) {
		defer wg.Done()
		z.Odd(func(z int) { fmt.Println(z) })
	}(z)
	go func(z *ZeroEvenOdd) {
		defer wg.Done()
		z.Even(func(z int) { fmt.Println(z) })
	}(z)

	wg.Wait()
}
