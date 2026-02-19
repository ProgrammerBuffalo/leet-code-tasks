package main

import (
	"fmt"
	"sync"
)

type H2O struct {
	ch1 chan struct{}
	ch2 chan struct{}
	ch3 chan struct{}
}

func NewH2O() *H2O {
    ch2 := make(chan struct{}, 1)
    ch2 <- struct{}{}
	return &H2O{
		ch1: make(chan struct{}, 1),
		ch2: ch2,
		ch3: make(chan struct{}, 1),
	}
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	select {
	case h.ch1 <- struct{}{}:
		releaseHydrogen()
        <-h.ch2
		return
	case h.ch2 <- struct{}{}:
		releaseHydrogen()
		h.ch3 <- struct{}{}
		return
	}
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	<-h.ch3
	releaseOxygen()
	<-h.ch1
}

func main() {
	h2o := NewH2O()

	wg := &sync.WaitGroup{}
	for range 4 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			h2o.Hydrogen(func() { fmt.Print("H") })
		}()
	}

	for range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			h2o.Oxygen(func() { fmt.Print("O") })
		}()
	}
	wg.Wait()
}
