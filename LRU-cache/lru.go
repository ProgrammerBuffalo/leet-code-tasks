package main

import (
	"hash/fnv"
)

type LRU struct {
	Buckets []*LinkedList
	Header  *Node

	Capacity int
}

func NewLRU(cap int) *LRU {
	return &LRU{
		Capacity: cap,
		Buckets:  make([]*LinkedList, cap),
	}
}

func (lru *LRU) Put(key string, val int) {
	list := lru.FindBucket(key)

	if list == nil {
		list = &LinkedList{}
		list.Add(key, val)
	}
}

func (lru *LRU) Get(key string) (int, bool) {
	list := lru.FindBucket(key)

	if list != nil {
		for node := list.Head; node != nil; node = node.Next {
			if node.Key == key {
				return node.Value, true
			}
		}
	}

	return 0, false
}

func (lru *LRU) Remove(key string) bool {
	list := lru.FindBucket(key)

	if list == nil {
		return false
	}

	ok := list.Remove(key)

	return ok
}

func (lru *LRU) FindBucket(key string) *LinkedList {
	index := generateHash(key) % uint32(lru.Capacity)
	list := lru.Buckets[index]

	return list
}

func generateHash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
