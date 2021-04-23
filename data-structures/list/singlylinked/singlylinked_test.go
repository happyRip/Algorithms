package singlylinked

import (
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	l := NewList()

	t.Run("Test NewNode()", func(t *testing.T) {
		want := node{1, nil}
		got := NewNode(1)

		if *got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Testing NewList()", func(t *testing.T) {
		var want list
		got := NewList()

		if *got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test PushBack()", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			l.PushBack(i)
		}

		want := []int{0, 1, 2, 3, 4}
		var got []int

		current := l.head
		for current.next != nil {
			got = append(got, current.value)
			current = current.next
		}
		got = append(got, current.value)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test PushFront()", func(t *testing.T) {
		l.PushFront(-1)
		l.PushFront(-2)

		want := []int{-2, -1, 0, 1, 2, 3, 4}
		var got []int

		current := l.head
		for current.next != nil {
			got = append(got, current.value)
			current = current.next
		}
		got = append(got, current.value)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test PopBack()", func(t *testing.T) {
		want := 4
		got, _ := l.PopBack()

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test PopFront()", func(t *testing.T) {
		want := -2
		got, _ := l.PopFront()

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Size()", func(t *testing.T) {
		want := 5
		got := l.Size()

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Insert()", func(t *testing.T) {
		l.Insert(5, 5)
		l.Insert(3, 3)
		l.Insert(0, 0)

		want := []int{0, -1, 0, 1, 3, 2, 3, 5}
		var got []int

		current := l.head
		for current.next != nil {
			got = append(got, current.value)
			current = current.next
		}
		got = append(got, current.value)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Remove()", func(t *testing.T) {
		l.Remove(0)
		l.Remove(3)
		l.Remove(5)

		want := []int{-1, 0, 1, 2, 3}
		var got []int

		current := l.head
		for current.next != nil {
			got = append(got, current.value)
			current = current.next
		}
		got = append(got, current.value)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test PeekBack()", func(t *testing.T) {
		want := 3
		got, _ := l.PeekBack()

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Testing PeekFront()", func(t *testing.T) {
		want := -1
		got, _ := l.PeekFront()

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Seek()", func(t *testing.T) {
		want := 1
		got, _ := l.Seek(2)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
