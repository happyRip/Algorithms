package doublylinked

import (
	"errors"
	"fmt"
)

type node struct {
	value    int
	next     *node
	previous *node
}

func NewNode(v int) *node {
	return &node{v, nil, nil}
}

type list struct {
	length int
	head   *node
}

func NewList() *list {
	return &list{}
}

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
	n.previous = current
	l.length++
	return nil
}

func (l *list) PushFront(v int) error {
	n := NewNode(v)

	n.next = l.head
	l.head = n
	l.length++
	return nil
}

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
	n.previous = current
	current.next = n
	if n.next != nil {
		n.next.previous = n
	}
	l.length++
	return nil
}

func (l *list) PopBack() (int, error) {
	if l.head == nil {
		return 0, errors.New("list is empty")
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	v := current.next.value
	current.next = nil
	l.length--
	return v, nil
}

func (l *list) PopFront() (int, error) {
	if l.head == nil {
		return 0, errors.New("list is empty")
	}

	v := l.head.value
	l.head = l.head.next
	l.length--
	return v, nil
}

func (l *list) Length() int {
	return l.length
}

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

func (l *list) ShowReverse() error {
	if l.head == nil {
		return errors.New("list is empty")
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	for current.previous != nil {
		fmt.Print(current.value, " ")
		current = current.previous
	}
	fmt.Println(current.previous)
	return nil
}
