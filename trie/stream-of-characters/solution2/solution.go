package main

import "fmt"

/*
In my second solution, I put the words into the trie in reverse order.

This allows me to avoid "cutting" the queue and trying every possible suffix
one by one when a match is not found.

Instead, I start from the most recently added character and traverse the
queue backwards, following the reversed word in the trie one time
*/
type StreamChecker struct {
	trie *Trie

	suffixQueue CycleQueue
}

type CycleQueue struct {
	left  int
	right int

	queue []byte
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

type Trie struct {
	children map[byte]*Trie
	isEnd    bool
}

func Constructor(words []string) StreamChecker {
	root := &Trie{children: make(map[byte]*Trie), isEnd: false}
	maxLen := 0

	for _, word := range words {
		cursor := root
		for i := len(word) - 1; i >= 0; i-- {
			if foundNode, ok := cursor.children[word[i]]; !ok {
				node := &Trie{children: make(map[byte]*Trie), isEnd: false}
				cursor.children[word[i]] = node
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

func (this *StreamChecker) Query(letter byte) bool {
	cursor := this.trie
	this.suffixQueue.push(letter)

	if this.suffixQueue.left <= this.suffixQueue.right {
		for i := this.suffixQueue.right; i >= this.suffixQueue.left; i-- {
			if foundNode, ok := cursor.children[this.suffixQueue.queue[i]]; !ok {
				return false
			} else {
				cursor = foundNode
				if cursor.isEnd {
					return true
				}
			}
		}
	} else {
		for i := this.suffixQueue.right; this.suffixQueue.left != this.suffixQueue.right; {
			if foundNode, ok := cursor.children[this.suffixQueue.queue[i]]; !ok {
				return false
			} else {
				cursor = foundNode
				if cursor.isEnd {
					return true
				}
			}
			if i == 0 {
				i = len(this.suffixQueue.queue) - 1
				continue
			}
			i--
		}
	}

	return cursor.isEnd
}

func main() {
	streamChecker := Constructor([]string{"ab", "ba", "aaab", "abab", "baa"})

	fmt.Println(1, streamChecker.Query('a'), false)
	fmt.Println(2, streamChecker.Query('b'), true)
	fmt.Println(3, streamChecker.Query('a'), true)
	fmt.Println(4, streamChecker.Query('b'), true)
	fmt.Println(5, streamChecker.Query('a'), true)
	fmt.Println(6, streamChecker.Query('b'), true)
	fmt.Println(7, streamChecker.Query('b'), false)
	fmt.Println(8, streamChecker.Query('b'), false)
	fmt.Println(9, streamChecker.Query('a'), true)
	fmt.Println(10, streamChecker.Query('b'), true)
	fmt.Println(11, streamChecker.Query('a'), true)
	fmt.Println(12, streamChecker.Query('b'), true)
	fmt.Println(13, streamChecker.Query('b'), false)
	fmt.Println(14, streamChecker.Query('b'), false)
}
