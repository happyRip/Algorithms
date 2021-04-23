package singlylinked

import (
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	l := NewList()

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

	t.Run("Test Length()", func(t *testing.T) {
		want := 5
		got := l.Length()

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Insert()", func(t *testing.T) {
		l.Insert(5, 7)
		l.Insert(3, 3)

		want := []int{-2, -1, 0, 3, 1, 2, 3, 4, 5}
		var got []int

		current := l.head
		for current.next != nil {
			got = append(got, current.value)
			current = current.next
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
