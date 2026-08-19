package main

import "fmt"

type Trie struct {
	children map[rune]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{children: make(map[rune]*Trie), isEnd: false}
}

func (this *Trie) Insert(word string) {
	cursor := this
	for _, ch := range word {
		if foundNode, ok := cursor.children[ch]; !ok {
			node := &Trie{children: make(map[rune]*Trie), isEnd: false}
			cursor.children[ch] = node
			cursor = node
		} else {
			cursor = foundNode
		}
	}
	cursor.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cursor := this
	for _, ch := range word {
		if foundNode, ok := cursor.children[ch]; !ok {
			return false
		} else {
			cursor = foundNode
		}
	}
	return cursor.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cursor := this
	for _, ch := range prefix {
		if foundNode, ok := cursor.children[ch]; !ok {
			return false
		} else {
			cursor = foundNode
		}
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // return True
	fmt.Println(trie.Search("app"))     // return False
	fmt.Println(trie.StartsWith("app")) // return True
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // return True
}
