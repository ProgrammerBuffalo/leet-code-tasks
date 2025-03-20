package main

type LinkedList struct {
	Head *Node
	Root *Node
}

func (l *LinkedList) Add(key string, val int) {
	n := &Node{
		Key:   key,
		Value: val,
	}

	if l.Head == nil {
		l.Head = n
		return
	}

	l.Root.Next = n
	l.Root = n
}

func (l *LinkedList) Remove(key string) bool {
	n := l.Find(key)

	if n == nil {
		return false
	}

	switch n {
	case l.Head:
		l.Head = l.Head.Next
	case l.Root:
		l.Root = l.Root.Previous
	default:
		n.Previous.Next = n.Next
		n.Next.Previous = n.Previous
	}

	return true
}

func (l *LinkedList) Find(key string) *Node {
	for n := l.Head; n != nil; n = n.Next {
		if n.Key == key {
			return n
		}
	}
	return nil
}
