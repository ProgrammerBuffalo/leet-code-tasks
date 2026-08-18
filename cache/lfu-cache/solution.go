package main

import "fmt"

// LFU cache is implemented using two hash maps:
// 1. keyMap stores key-value pairs and provides O(1) access by key.
// 2. freqMap groups keys by their access frequency and provides O(1) frequency updates.
type LFUCache struct {
	keyMap  map[int]*Node
	freqMap map[int]*DoubleLinkedList

	capacity int

	minFreq int
	maxFreq int
}

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Val  int
	Key  int
	Freq int

	Next *Node
	Prev *Node
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		keyMap:   make(map[int]*Node),
		freqMap:  make(map[int]*DoubleLinkedList),
		capacity: capacity,
		minFreq:  1,
		maxFreq:  1,
	}
}

func (this *LFUCache) Get(key int) int {
	if n, ok := this.keyMap[key]; ok {
		n.Freq++
		if n.Freq > this.maxFreq {
			this.maxFreq = n.Freq
		}
		this.addToHigherFreq(n)
		return n.Val
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	n, ok := this.keyMap[key]
	if len(this.keyMap) == this.capacity && !ok {
		leastNode := this.removeLeastUsed()
		if leastNode != nil {
			delete(this.keyMap, leastNode.Key)
		}
	}
	if ok {
		n.Freq++
		n.Val = value
		if n.Freq > this.maxFreq {
			this.maxFreq = n.Freq
		}
		this.addToHigherFreq(n)
	} else {
		n = &Node{Val: value, Key: key, Freq: 1}
		this.minFreq = 1
		this.keyMap[key] = n
		this.addNew(n)
	}
}

func (this *LFUCache) removeLeastUsed() *Node {
	if freqDll, ok := this.freqMap[this.minFreq]; ok {
		leastNode := freqDll.Tail
		if freqDll.Tail != freqDll.Head {
			leastNode.Prev.Next = nil
			freqDll.Tail = leastNode.Prev
			leastNode.Next = nil
		} else {
			delete(this.freqMap, this.minFreq)
			if this.minFreq != 1 {
				for i := this.minFreq; i <= this.maxFreq; i++ {
					if _, ok = this.freqMap[i]; ok {
						this.minFreq = i
						break
					}
				}
			}
		}
		return leastNode
	}
	return nil
}

func (this *LFUCache) addNew(n *Node) {
	if freqDll, ok := this.freqMap[n.Freq]; ok {
		if freqDll.Head == freqDll.Tail {
			freqDll.Head = n
			n.Next = freqDll.Tail
			freqDll.Tail.Prev = n
		} else {
			freqDll.Head.Prev = n
			n.Next = freqDll.Head
			freqDll.Head = n
		}
	} else {
		this.freqMap[n.Freq] = &DoubleLinkedList{Head: n, Tail: n}
	}
}

func (this *LFUCache) addToHigherFreq(n *Node) {
	if freqDll, ok := this.freqMap[n.Freq-1]; ok {
		if n == freqDll.Head && n == freqDll.Tail {
			delete(this.freqMap, n.Freq-1)
			n.Prev = nil
			n.Next = nil
			this.addNew(n)
			if this.minFreq == n.Freq-1 {
				this.minFreq = this.maxFreq
				for i := n.Freq - 1; i <= this.maxFreq; i++ {
					if _, ok = this.freqMap[i]; ok {
						this.minFreq = i
						break
					}
				}
			}
		} else {
			if n != freqDll.Head {
				n.Prev.Next = n.Next
			} else {
				freqDll.Head = freqDll.Head.Next
				freqDll.Head.Prev = nil
			}
			if n != freqDll.Tail {
				n.Next.Prev = n.Prev
			} else {
				freqDll.Tail = freqDll.Tail.Prev
				freqDll.Tail.Next = nil
			}
			n.Prev = nil
			n.Next = nil
			this.addNew(n)
		}
	}
}

func main() {
	lfu := Constructor(2)

	lfu.Put(1, 1)
	lfu.Put(1, 1)
	lfu.Put(1, 1)

	lfu.Put(2, 2)
	lfu.Put(2, 2)

	lfu.Put(3, 3)

	fmt.Println(lfu.Get(1))
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(3))
}
