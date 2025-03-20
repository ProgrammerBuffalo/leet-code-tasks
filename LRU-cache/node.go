package main

type Node struct {
	Key      string
	Value    int
	Next     *Node
	Previous *Node
}
