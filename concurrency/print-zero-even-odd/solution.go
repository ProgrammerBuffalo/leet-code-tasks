package main

import (
	"fmt"
	"sync"
)

type ZeroEvenOdd struct {
	n     int
	currN int

	evenCh chan struct{}
	oddCh  chan struct{}
	zeroCh chan struct{}
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeroCh := make(chan struct{}, 1)

	return &ZeroEvenOdd{
		n:     n,
		currN: 0,

		evenCh: make(chan struct{}, 1),
		oddCh:  make(chan struct{}, 1),
		zeroCh: zeroCh,
	}
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for range z.zeroCh {
		printNumber(z.currN)
		if z.currN%2 == 0 {
			z.evenCh <- struct{}{}
		} else {
			z.oddCh <- struct{}{}
		}
	}
	close(z.oddCh)
	close(z.evenCh)
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for range z.evenCh {
		printNumber(z.currN)
		z.currN++
		if z.currN == z.n {
			close(z.zeroCh)
		}
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for range z.oddCh {
		printNumber(z.currN)
		z.currN++
		if z.currN == z.n {
			close(z.zeroCh)
		}
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
