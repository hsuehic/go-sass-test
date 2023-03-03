package main

import "fmt"

type Node[T comparable] struct {
	next *Node[T]
	val  T
}

type LinkedList[T comparable] struct {
	head *Node[T]
	size int
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) InsertAt(v T, index int) {
	if index > l.Size() || index < 0 {
		panic("Index out of range")
	}
	newNode := &Node[T]{next: nil, val: v}
	if index == 0 {

	} else {
		current := l.head
		for i := 0; i < index-1; {
			i++
			current = current.next
		}
		newNode.next = current.next
		current.next = newNode
	}
	l.size++
}

func (l *LinkedList[T]) InsertAfter(v, n *Node[T]) {

}

func (l *LinkedList[T]) Shift() T {
	var v T
	if l.Size() != 0 {
		v = l.head.val
		l.head = l.head.next
		l.size--
	}
	return v
}

func (l *LinkedList[T]) Unshift(v T) {

}

func (l *LinkedList[T]) Append(v T) {
	newNode := &Node[T]{nil, v}

	if l.Size() == 0 {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode

	}
	l.size++
}

func (l *LinkedList[T]) Pop() T {
	var v T
	if l.Size() == 0 {
		panic("Indx out of Range")
	} else if l.Size() == 1 {
		v = l.head.val
		l.head = nil
		l.size = 0
	} else {
		current := l.head
		for current.next != nil {
			node := current.next
			if current.next.next == nil {
				current.next = nil
				v = node.val
			}
			current = node
		}

		l.size--
	}
	return v
}

func (l *LinkedList[T]) Remove(v T) bool {
	var pre, cur, next *Node[T]
	var flag bool = false
	pre = nil
	cur = l.head
	next = cur.next
	for cur.next != nil {
		if cur.val == v {
			flag = true
			break
		}
		pre = cur
		cur = next
		next = next.next
	}
	if flag {
		if pre != nil {
			pre.next = next
		}
		l.size--
	} else {
		panic("Node doesn't exist in the list")
	}
	return flag
}

func (l *LinkedList[T]) RemoveNode(n *Node[T]) bool {
	v := n.val
	return l.Remove(v)
}

func main() {
	l := &LinkedList[int]{nil, 0}
	for i := 0; i < 10; {
		l.Append(i)
		i++
	}
	node := l.head
	for node != nil {
		fmt.Println(node.val)
		node = node.next
	}
	fmt.Printf("List has %v nodes\n", l.Size())
}
