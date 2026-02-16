package main

import (
	"fmt"
	"sync"
)

type Foo struct {
	ch1 chan struct{}
	ch2 chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		ch1: make(chan struct{}),
		ch2: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	printFirst()
	f.ch1 <- struct{}{}
}

func (f *Foo) Second(printSecond func()) {
	<-f.ch1
	printSecond()
	f.ch2 <- struct{}{}
}

func (f *Foo) Third(printThird func()) {
	<-f.ch2
	printThird()
}

func main() {
	f := NewFoo()

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go f.Third(func() {
		fmt.Print("third ")
		defer wg.Done()
	})
	go f.Second(func() {
		fmt.Print("second ")
		defer wg.Done()
	})
	go f.First(func() {
		fmt.Print("first ")
		defer wg.Done()
	})

	wg.Wait()
}
