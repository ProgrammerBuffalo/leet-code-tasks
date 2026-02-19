package main

import (
	"fmt"
	"sync"
)

/*
The output array describes the calls you made to the functions controlling the forks and the eat function, its format is:
output[i] = [a, b, c] (three integers)
- a is the id of a philosopher.
- b specifies the fork: {1 : left, 2 : right}.
- c specifies the operation: {1 : pick, 2 : put, 3 : eat}.
*/

type DiningPhilosophers struct {
	forks []chan struct{}
	once  sync.Once
}

func (this *DiningPhilosophers) wantsToEat(
	philosopher int,
	pickLeftFork func(),
	pickRightFork func(),
	eat func(),
	putLeftFork func(),
	putRightFork func(),
) {
	this.once.Do(func() {
		this.forks = make([]chan struct{}, 5)
		for i := range 5 {
			this.forks[i] = make(chan struct{}, 1)
		}
	})
	left, right := philosopher%5, (philosopher+1)%5
	if philosopher%2 == 0 {
		this.forks[right] <- struct{}{}
		this.forks[left] <- struct{}{}
		pickRightFork()
		pickLeftFork()
	} else {
		this.forks[left] <- struct{}{}
		this.forks[right] <- struct{}{}
		pickLeftFork()
		pickRightFork()
	}
	eat()
	<-this.forks[left]
	<-this.forks[right]
	putLeftFork()
	putRightFork()
}

func main() {
	dp := &DiningPhilosophers{}
	wg := &sync.WaitGroup{}
	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dp.wantsToEat(i,
				func() {
					fmt.Printf("[%d,%d,1]", i, i%5)
				},
				func() {
					fmt.Printf("[%d,%d,1]", i, (i+1)%5)
				},
				func() {
					fmt.Printf("[%d,0,3]", i)
				},
				func() {
					fmt.Printf("[%d,%d,2]", i, i%5)
				},
				func() {
					fmt.Printf("[%d,%d,2]", i, (i+1)%5)
				})
		}()
	}
	wg.Wait()
}
