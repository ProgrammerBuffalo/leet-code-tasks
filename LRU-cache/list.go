package main

type LinkedList struct {
	Head *Node
	Root *Node
}

func (l *LinkedList) Add(key string, val int) *Node {
	n := &Node{
		Key:   key,
		Value: val,
	}

	if l.Head == nil {
		l.Head = n
		return n
	}

	l.Root.Next = n
	l.Root = n

	return n
}

func (l *LinkedList) RemoveByKey(key string) *Node {
	node := l.FindByKey(key)

	if node == nil {
		return nil
	}

	l.ChangeRefs(node)

	return node
}

func (l *LinkedList) RemoveByNode(node *Node) *Node {
	ok := l.IsNodeExists(node)

	if ok == false {
		return nil
	}

	l.ChangeRefs(node)

	return node
}

func (l *LinkedList) ChangeRefs(node *Node) {
	switch node {
	case l.Head:
		l.Head = l.Head.Next
	case l.Root:
		l.Root = l.Root.Previous
	default:
		node.Previous.Next = node.Next
		node.Next.Previous = node.Previous
	}
}

func (l *LinkedList) FindByKey(key string) *Node {
	for n := l.Head; n != nil; n = n.Next {
		if n.Key == key {
			return n
		}
	}
	return nil
}

func (l *LinkedList) IsNodeExists(node *Node) bool {
	for n := l.Head; n != nil; n = n.Next {
		if node == n {
			return true
		}
	}
	return false
}
