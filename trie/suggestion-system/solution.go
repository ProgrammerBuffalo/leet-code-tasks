package main

import (
	"fmt"
	"strings"
)

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) AddProducts(products []string) {
	var child *Trie
	for _, p := range products {
		child = t
		for i := 0; i < len(p)-1; i++ {
			if child.children[p[i]-'a'] == nil {
				child.children[p[i]-'a'] = &Trie{children: [26]*Trie{}}
			}
			child = child.children[p[i]-'a']
		}
		if child.children[p[len(p)-1]-'a'] == nil {
			child.children[p[len(p)-1]-'a'] = &Trie{children: [26]*Trie{}}
		}
		child.children[p[len(p)-1]-'a'].isEnd = true
	}
}

/*
'a' - is 62 byte 'b' - is 63 byte,
so it means that if I make letter - 'a' I ll get exact position of array
26 - is eng alphabet length
*/
func main() {
	fmt.Println(suggestedProducts([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"))
}

func suggestedProducts(products []string, searchWord string) [][]string {
	root := &Trie{children: [26]*Trie{}}
	root.AddProducts(products)

	var prefix strings.Builder
	res := make([][]string, 0, len(searchWord))
	for i := 0; i < len(searchWord); i++ {
		prefix.WriteByte(searchWord[i])
		res = append(res, suggestByPrefix(root, prefix.String()))
	}
	return res
}

func suggestByPrefix(root *Trie, prefix string) []string {
	it := root
	fromSuggestion := make([]byte, 0)
	for i := 0; i < len(prefix); i++ {
		if it.children[prefix[i]-'a'] == nil {
			return nil
		}
		it = it.children[prefix[i]-'a']
	}
	fromSuggestion = append(fromSuggestion, prefix...)
	suggestions := make([]string, 0, 3)

	if it.isEnd {
		suggestions = append(suggestions, string(fromSuggestion))
	}

	var dfs func(curr *Trie)
	dfs = func(curr *Trie) {
		for i := 0; i < 26; i++ {
			if len(suggestions) == 3 {
				return
			}
			if curr.children[i] != nil {
				fromSuggestion = append(fromSuggestion, byte(i)+'a')
				if curr.children[i].isEnd {
					suggestions = append(suggestions, string(fromSuggestion))
				}
				dfs(curr.children[i])
			}
		}
		if len(fromSuggestion) > 0 {
			fromSuggestion = fromSuggestion[:len(fromSuggestion)-1]
		} else {
			return
		}
	}
	dfs(it)

	return suggestions
}
