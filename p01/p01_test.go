package p01

import (
	"container/list"
  "testing"
)

func TestLastElement(t *testing.T) {
	l := list.New()

	lastElement := 99

	l.PushBack(1)
	l.PushBack(2)
  l.PushBack(lastElement)

	result := LastElement(l)

  if result.Value != lastElement {
    t.Errorf("Result/Expected: %d/%d", result.Value, lastElement)
	}
}
