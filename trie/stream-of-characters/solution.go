package main

import "fmt"

type StreamChecker struct {
	trie *Trie

	suffixQueue CycleQueue
}

type CycleQueue struct {
	left  int
	right int

	queue []byte
}

type Trie struct {
	children map[byte]*Trie
	isEnd    bool
}

func Constructor(words []string) StreamChecker {
	root := &Trie{children: make(map[byte]*Trie), isEnd: false}
	maxLen := 0
	for _, word := range words {
		cursor := root
		for _, ch := range word {
			if foundNode, ok := cursor.children[byte(ch)]; !ok {
				node := &Trie{children: make(map[byte]*Trie), isEnd: false}
				cursor.children[byte(ch)] = node
				cursor = node
			} else {
				cursor = foundNode
			}
		}
		cursor.isEnd = true

		if len(word) > maxLen {
			maxLen = len(word)
		}
	}

	return StreamChecker{
		trie: root,
		suffixQueue: CycleQueue{
			left:  0,
			right: 0,
			queue: make([]byte, maxLen),
		}}
}

func (cq *CycleQueue) push(letter byte) {
	if cq.queue[0] == 0 {
		cq.queue[cq.right] = letter
		return
	}
	cq.right++
	if cq.right == len(cq.queue) {
		cq.right = 0
		cq.left = 1
	}
	cq.queue[cq.right] = letter
	if cq.left == cq.right {
		cq.left++
		if cq.left == len(cq.queue) {
			cq.left = 0
		}
	}
}

func (this *StreamChecker) Query(letter byte) bool {
	cursor := this.trie
	this.suffixQueue.push(letter)
	for i := this.suffixQueue.left; ; i = (i + 1) % len(this.suffixQueue.queue) {
		if foundNode, ok := cursor.children[this.suffixQueue.queue[i]]; ok {
			cursor = foundNode
			if cursor.isEnd && i == this.suffixQueue.right {
				return true
			}
		}
		if i == this.suffixQueue.right {
			break
		}
	}
	return false
}

func main() {
	streamChecker := Constructor([]string{"cd", "f", "kl"})

	fmt.Println(streamChecker.Query('a'))
	fmt.Println(streamChecker.Query('b'))
	fmt.Println(streamChecker.Query('c'))
	fmt.Println(streamChecker.Query('d'))
	fmt.Println(streamChecker.Query('e'))
	fmt.Println(streamChecker.Query('f'))
	fmt.Println(streamChecker.Query('g'))
	fmt.Println(streamChecker.Query('h'))
	fmt.Println(streamChecker.Query('i'))
	fmt.Println(streamChecker.Query('j'))
	fmt.Println(streamChecker.Query('k'))
	fmt.Println(streamChecker.Query('l'))
}
