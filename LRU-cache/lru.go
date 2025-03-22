package main

import (
	"hash/fnv"
)

type LRU struct {
	Buckets []*LinkedList
	Header  *Node

	Capacity  int
	NodeCount int
}

func NewLRU(cap int) *LRU {
	return &LRU{
		Capacity:  cap,
		NodeCount: 0,
		Buckets:   make([]*LinkedList, cap),
		Header: &Node{
			Value: 0,
		},
	}
}

func (lru *LRU) Put(key string, val int) {
	list := lru.FindBucket(key)

	if list == nil {
		list = &LinkedList{}
	}

	node := list.Add(key, val)

	node.After = lru.Header
	lru.Header.Before = node

	lru.NodeCount++

	if lru.Header.After == nil {
		lru.Header.After = node
	}

	if lru.NodeCount == lru.Capacity {
		n := list.RemoveByNode(lru.Header.After)
		lru.Header.After = n.After

		n.After.Before = lru.Header.Before

		n.After = nil
		n.Before = nil
	}
}

func (lru *LRU) Get(key string) (int, bool) {
	list := lru.FindBucket(key)

	if list != nil {
		for node := list.Head; node != nil; node = node.Next {
			if node.Key == key {
				lru.Header.Before = node
				node.After = lru.Header
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

	node := list.RemoveByKey(key)

	if node == nil {
		return false
	}

	node.Before.After = node.After
	node.After.Before = node.Before

	lru.NodeCount--

	return true
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
