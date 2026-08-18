package main

import (
	"fmt"
	"lru-cache/lru"
)

func main() {
	l := lru.LRU{}
	fmt.Print(l)
}
