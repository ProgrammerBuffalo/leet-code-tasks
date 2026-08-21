package main

import "fmt"

/*
In my first solution, I tried to put the words into the trie exactly as they
are presented in the input.

I used a circular queue for memory efficiency. Its size is equal to the
length of the longest word in words, because it is impossible for a valid
suffix to be longer than the longest word in the trie.

However, my first approach had a problem with the number of iterations.
For example, suppose the circular queue contains:

a b c d

If I start traversing the trie from `a` and cannot find the next node, I have
to try again without the first character. The possible suffix could be `bcd`,
which may exist in the trie.

So I would have to try:

a -> b -> c -> d
b -> c -> d
c -> d
d

It results in O(N²) time complexity.

This causes the number of iterations to grow very quickly when the stream
contains many characters or the words are long.

In the test case 19/20 i got TLE
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
	for i, j := this.suffixQueue.left, this.suffixQueue.left; ; i = (i + 1) % len(this.suffixQueue.queue) {
		if foundNode, ok := cursor.children[this.suffixQueue.queue[i]]; ok {
			cursor = foundNode
			if cursor.isEnd && i == this.suffixQueue.right {
				return true
			}
		} else {
			cursor = this.trie
			i = j
			j = (j + 1) % len(this.suffixQueue.queue)
		}
		if j == this.suffixQueue.right {
			if cursor == this.trie {
				rootChild := cursor.children[letter]
				if rootChild != nil {
					return rootChild.isEnd
				}
			}
			return false
		}
	}
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
