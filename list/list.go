package list

import (
	"errors"
	"fmt"
)

// Node ...
type Node struct {
	Value int
	Next  *Node
}

// List ...
type List struct {
	Head *Node
}

// PushBack inserts element at the end of the List.
func (l *List) PushBack(v int) error {
	node := &Node{
		Value: v,
	}

	if l.Head == nil {
		l.Head = node
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
	return nil
}

// PushFront inserts element at the beginning of the List.
func (l *List) PushFront(v int) error {
	node := &Node{
		Value: v,
	}

	if l.Head == nil {
		l.Head = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
	return nil
}

// Insert adds Node to the List on selected position.
func (l *List) Insert(v, pos int) error {
	if pos < 0 {
		return errors.New("position out of scope")
	}

	node := &Node{
		Value: v,
	}

	if pos == 0 {
		node.Next = l.Head
		l.Head = node

		return nil
	}

	current := l.Head
	for i := 0; i < pos-1; i++ {
		if current.Next == nil {
			return errors.New("position out of scope")
		}

		current = current.Next
	}
	node.Next = current.Next
	current.Next = node

	return nil
}

// PopBack removes element from the end of the List and returns it's value.
func (l *List) PopBack() (int, error) {
	if l.Head == nil {
		return 0, errors.New("list is empty")
	}

	current := l.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	v := current.Next.Value
	current.Next = nil

	return v, nil
}

// PopFront removes element from the front of the List and returns it's value.
func (l *List) PopFront() (int, error) {
	if l.Head == nil {
		return 0, errors.New("list is empty")
	}

	v := l.Head.Value

	if l.Head.Next != nil {
		l.Head = l.Head.Next
	} else {
		l.Head = nil
	}

	return v, nil
}

// Remove deletes an element on selected position from the List.
func (l *List) Remove(pos int) error {
	if pos < 0 {
		return errors.New("position out of scope")
	} else if pos == 0 {
		l.Head = l.Head.Next

		return nil
	}

	current := l.Head
	for i := 0; i < pos-1; i++ {
		if current.Next == nil {
			return errors.New("position out of scope")
		}
		current = current.Next
	}
	current.Next = current.Next.Next

	return nil
}

// Show is used to print the List.
func (l *List) Show() {
	current := l.Head

	for current.Next != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println(current.Value)
}
