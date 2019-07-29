package linkedlist

import (
	"errors"
	"fmt"
)

// ErrEmptyList indicates the List is empty
var ErrEmptyList = errors.New("empty list")

// Node represents nodes used in a double linked list
// Val can be of any type
// next is a pointer to the next node in a list
// prev is a pointer to the previous node in a list
type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

// Next pointer to the next node.
func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	return n.next
}

// Prev pointer to the previous node.
func (n *Node) Prev() *Node {
	if n == nil {
		return nil
	}
	return n.prev
}

// First pointer to the first node (head).
func (n *Node) First() *Node {
	if n.prev == nil {
		return n
	}
	return n.prev.Prev()
}

// Last pointer to the last node (tail).
func (n *Node) Last() *Node {
	if n.next == nil {
		return n
	}
	return n.next.Last()
}

// List represents a Linked List that holds references to the First and Last Nodes
type List struct {
	first *Node
	last  *Node
}

// NewList creates a new linked list based on values provided as an argument
func NewList(values ...interface{}) *List {
	list := &List{}
	for _, v := range values {
		list.PushBack(v)
	}
	return list
}

// First()`: pointer to the first node (head).
func (l *List) First() *Node {
	return l.first
}

// Last()`: pointer to the last node (tail).
func (l *List) Last() *Node {
	return l.last
}

// PushBack(v interface{})`: insert value at back.
func (l *List) PushBack(v interface{}) {
	originalLast := l.Last()
	newLast := &Node{
		Val: v,
	}
	l.last = newLast
	if originalLast != nil {
		newLast.prev = originalLast
		originalLast.next = newLast
	} else if l.First() == nil && originalLast == nil {
		// if the list was empty then the last node is the first one too
		l.first = newLast

	}
}

// PopBack() (interface{}, error): remove value at back.
func (l *List) PopBack() (interface{}, error) {
	last := l.Last()
	if last == nil {
		return nil, ErrEmptyList
	}
	l.last = last.Prev()
	if l.last != nil {
		l.last.next = nil
	} else {
		// clear first if last is nil
		l.first = nil
	}
	return last.Val, nil
}

// PushFront(v interface{}) `: insert value at front.
func (l *List) PushFront(v interface{}) {
	origianlFirst := l.First()
	newFirst := &Node{
		Val: v,
	}

	l.first = newFirst
	if origianlFirst != nil {
		newFirst.next = origianlFirst
		origianlFirst.prev = newFirst
	} else if l.Last() == nil && origianlFirst == nil {
		// if the list was empty then the first node is the last one too
		l.last = newFirst
	}
}

// PopFront() (interface{}, error)`: remove value at front.
func (l *List) PopFront() (interface{}, error) {
	first := l.First()
	if first == nil {
		return nil, ErrEmptyList
	}
	l.first = first.Next()
	if l.first != nil {
		l.first.prev = nil
	} else {
		// clear last just in case
		l.last = nil
	}
	return first.Val, nil
}

// Reverse()`: reverse the linked list.
func (l *List) Reverse() {
	current := l.First()
	for current != nil {
		newPrev := current.Next()
		current.next = current.Prev()
		current.prev = newPrev
		if current.Prev() == nil {
			l.first = current
		}
		if current.Next() == nil {
			l.last = current
		}
		current = current.Prev()
	}
}

func (l *List) Traverse() {
	l.traverse()
}
func (l *List) traverse() {
	n := l.First()
	for n != nil {
		fmt.Printf("%+v\n", n.Val)
		n = n.Next()
	}
}
