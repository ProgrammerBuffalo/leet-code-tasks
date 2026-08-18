package lru

import (
	"hash/fnv"
	"lru-cache/list"
	"lru-cache/node"
)

type LRU struct {
	Buckets []*list.LinkedList
	Header  *node.Node

	Capacity  int
	NodeCount int
}

func NewLRU(cap int) *LRU {
	return &LRU{
		Capacity:  cap,
		NodeCount: 0,
		Buckets:   make([]*list.LinkedList, cap),
		Header: &node.Node{
			Value: 0,
		},
	}
}

func (lru *LRU) Put(key string, val int) {
	l := lru.FindBucket(key)

	if l == nil {
		l = &list.LinkedList{}
	}

	n := l.Add(key, val)

	n.After = lru.Header
	lru.Header.Before = n

	lru.NodeCount++

	if lru.Header.After == nil {
		lru.Header.After = n
	}

	if lru.NodeCount == lru.Capacity {
		n := l.RemoveByNode(lru.Header.After)
		lru.Header.After = n.After

		n.After.Before = lru.Header.Before

		n.After = nil
		n.Before = nil
	}
}

func (lru *LRU) Get(key string) (int, bool) {
	l := lru.FindBucket(key)

	if l != nil {
		for n := l.Head; n != nil; n = n.Next {
			if n.Key == key {
				lru.Header.Before = n
				n.After = lru.Header
				return n.Value, true
			}
		}
	}

	return 0, false
}

func (lru *LRU) Remove(key string) bool {
	l := lru.FindBucket(key)

	if l == nil {
		return false
	}

	n := l.RemoveByKey(key)

	if n == nil {
		return false
	}

	n.Before.After = n.After
	n.After.Before = n.Before

	lru.NodeCount--

	return true
}

func (lru *LRU) FindBucket(key string) *list.LinkedList {
	index := generateHash(key) % uint32(lru.Capacity)
	return lru.Buckets[index]
}

func generateHash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
