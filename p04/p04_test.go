package p04

import (
	"container/list"
  "testing"
)

func TestListLength(t *testing.T) {
	l := list.New()

	l.PushBack(1)
  l.PushBack(2)
	l.PushBack(3)

	result := ListLength(l)
	expected := l.Len()

  if result != expected {
    t.Errorf("Result/Expected: %d/%d", result, expected)
	}
}
