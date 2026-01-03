package main

import (
	"fmt"
	"math"
)

type Node struct {
	val int
	min int
}

type MinStack struct {
	stack []Node
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]Node, 0),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, Node{
			val: val,
			min: val,
		})
		return
	}

	min := this.stack[len(this.stack)-1].min

	if min > val {
		min = val
	}

	this.stack = append(this.stack, Node{
		val: val,
		min: min,
	})
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}

	this.stack = this.stack[0 : len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return math.MinInt32
	}

	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	if len(this.stack) == 0 {
		return math.MinInt32
	}

	return this.stack[len(this.stack)-1].min
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}
