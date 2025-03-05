package main

import (
	"fmt"
)

type FooBar struct {
	n    int
	foo  chan struct{}
	bar  chan struct{}
	done chan struct{}
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		n:    n,
		foo:  make(chan struct{}),
		bar:  make(chan struct{}),
		done: make(chan struct{})}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		fb.foo <- struct{}{}
		printFoo()
		<-fb.bar
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.foo
		printBar()
		fb.bar <- struct{}{}
	}

	close(fb.done)
}

func main() {
	f := NewFooBar(10)

	go f.Foo(func() { fmt.Print("foo") })
	go f.Bar(func() { fmt.Print("bar") })

	select {
	case <-f.done:
		{
			return
		}
	}
}
