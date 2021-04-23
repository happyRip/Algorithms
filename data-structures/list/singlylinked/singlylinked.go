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
	head *node
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
		return nil
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = n
	return nil
}

// PushFront inserts element at the beginning of the list.
func (l *list) PushFront(v int) error {
	n := NewNode(v)

	n.next = l.head
	l.head = n
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
	return nil
}

// PeekBack returns the last value of a list without removing it
func (l *list) PeekBack() (int, error) {
	if l.head == nil {
		return 0, errors.New("list is empty")
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	return current.value, nil
}

// PeekFront returns the first value of a list without removing it
func (l *list) PeekFront() (int, error) {
	if l.head == nil {
		return 0, errors.New("list is empty")
	}
	return l.head.value, nil
}

// Seek returns the value located on a position pos without removing it
func (l *list) Seek(pos int) (int, error) {
	if pos < 0 {
		return 0, errors.New("position out of scope")
	} else if l.head == nil {
		return 0, errors.New("list is empty")
	} else if pos == 0 {
		return l.PeekFront()
	}

	current := l.head
	for i := 0; i < pos-1; i++ {
		if current.next == nil {
			return 0, errors.New("position out of scope")
		}
		current = current.next
	}
	return current.next.value, nil
}

// Size is used to get size of the list
func (l *list) Size() int {
	if l.head == nil {
		return 0
	}

	var size int
	current := l.head
	for current.next != nil {
		size++
		current = current.next
	}
	size++
	return size
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
