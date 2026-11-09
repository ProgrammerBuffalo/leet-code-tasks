package main

import "fmt"

type SmallestInfiniteSet struct {
	cursorInSet    int
	addedBackItems []int
	set            map[int]struct{}
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		cursorInSet:    1,
		addedBackItems: make([]int, 0),
		set:            make(map[int]struct{}),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(this.addedBackItems) == 0 {
		this.cursorInSet++
		return this.cursorInSet - 1
	}

	popped := this.addedBackItems[0]
	delete(this.set, popped)

	if len(this.addedBackItems) == 1 {
		this.addedBackItems = this.addedBackItems[1:]
		return popped
	}

	this.addedBackItems[0] = this.addedBackItems[len(this.addedBackItems)-1]
	this.addedBackItems = this.addedBackItems[:len(this.addedBackItems)-1]

	if len(this.addedBackItems) == 1 {
		return popped
	}

	for i, l, r := 0, 1, 2; ; {
		l = 2*i + 1
		r = 2*i + 2
		if r >= len(this.addedBackItems) {
			if l < len(this.addedBackItems) && this.addedBackItems[l] < this.addedBackItems[i] {
				this.addedBackItems[i], this.addedBackItems[l] = this.addedBackItems[l], this.addedBackItems[i]
			}
			return popped
		}
		if this.addedBackItems[l] > this.addedBackItems[r] {
			if this.addedBackItems[r] < this.addedBackItems[i] {
				this.addedBackItems[i], this.addedBackItems[r] = this.addedBackItems[r], this.addedBackItems[i]
				i = r
			} else {
				return popped
			}
		} else {
			if this.addedBackItems[l] < this.addedBackItems[i] {
				this.addedBackItems[i], this.addedBackItems[l] = this.addedBackItems[l], this.addedBackItems[i]
				i = l
			} else {
				return popped
			}
		}
	}
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if this.cursorInSet <= num {
		return
	}

	if _, ok := this.set[num]; ok {
		return
	}

	this.set[num] = struct{}{}
	this.addedBackItems = append(this.addedBackItems, num)

	if len(this.addedBackItems) == 1 {
		return
	}

	for i, p := len(this.addedBackItems)-1, 0; i >= 0; {
		p = (i - 1) / 2
		if this.addedBackItems[p] > this.addedBackItems[i] {
			this.addedBackItems[p], this.addedBackItems[i] = this.addedBackItems[i], this.addedBackItems[p]
			i = p
		} else {
			return
		}
	}
}

func main() {
	s := Constructor()

	s.PopSmallest() // 1
	s.PopSmallest() // 2
	s.PopSmallest() // 3
	s.PopSmallest() // 4

	s.AddBack(1)
	s.AddBack(3)
	s.AddBack(2)
	s.AddBack(5)
	s.AddBack(4)

	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
}

//     1
//   5   3
//  6 7 8 9
// 3
