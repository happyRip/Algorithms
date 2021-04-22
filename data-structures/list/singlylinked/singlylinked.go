package singlylinked

import (
	"errors"
	"fmt"
)

type node struct {
	value int
	next  *node
}

// NewList returns a new node instance
func NewNode(v int) *node {
	return &node{v, nil}
}

type list struct {
	length int
	head   *node
}

// NewList returns a new list list instance
func NewList() *list {
	return &list{}
}

// PushBack inserts element at the end of the list.
func (l *list) PushBack(v int) error {
	n := NewNode(v)

	if l.head == nil {
		l.head = n
		l.length++
		return nil
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = n
	l.length++
	return nil
}

// PushFront inserts element at the beginning of the list.
func (l *list) PushFront(v int) error {
	n := NewNode(v)

	n.next = l.head
	l.head = n
	l.length++
	return nil
}

// Insert adds node to the list on selected position.
func (l *list) Insert(v, pos int) error {
	if pos < 0 {
		return errors.New("position out of scope")
	}

	if pos == 0 {
		return l.PushFront(v)
	}

	n := NewNode(v)

	current := l.head
	for i := 0; i < pos-1; i++ {
		if current.next == nil {
			return errors.New("position out of scope")
		}
		current = current.next
	}
	n.next = current.next
	current.next = n
	l.length++
	return nil
}

// PopBack removes element from the end of the list and returns it's value.
func (l *list) PopBack() (int, error) {
	if l.head == nil {
		return 0, errors.New("list is empty")
	}

	current := l.head
	for current.next.next != nil {
		current = current.next
	}
	v := current.next.value
	current.next = nil
	l.length--
	return v, nil
}

// PopFront removes element from the front of the list and returns it's value.
func (l *list) PopFront() (int, error) {
	if l.head == nil {
		return 0, errors.New("list is empty")
	}

	v := l.head.value
	if l.head.next != nil {
		l.head = l.head.next
	} else {
		l.head = nil
	}
	l.length--
	return v, nil
}

// Remove deletes an element on selected position from the list.
func (l *list) Remove(pos int) error {
	if pos < 0 {
		return errors.New("position out of scope")
	} else if l.head == nil {
		return errors.New("list is empty")
	} else if pos == 0 {
		l.head = l.head.next
		l.length--
		return nil
	}

	current := l.head
	for i := 0; i < pos-1; i++ {
		if current.next == nil {
			return errors.New("position out of scope")
		}
		current = current.next
	}
	current.next = current.next.next
	l.length--
	return nil
}

func (l *list) Length() int {
	return l.length
}

// Show is used to print the list.
func (l *list) Show() error {
	if l.head == nil {
		return errors.New("list is empty")
	}

	current := l.head

	for current.next != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println(current.value)
	return nil
}
