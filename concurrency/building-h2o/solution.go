package main

import (
	"fmt"
	"sync"
)

type H2O struct {
	ch1 chan struct{}
	ch2 chan struct{}
}

func NewH2O() *H2O {
	return &H2O{}
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	releaseHydrogen()
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	releaseOxygen()
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
