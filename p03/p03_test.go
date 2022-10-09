package p03

import (
	"container/list"
  "testing"
)

func TestNthElement(t *testing.T) {
	l := list.New()

	nthElement := 99

	l.PushBack(1)
  l.PushBack(nthElement)
	l.PushBack(3)

	result := NthElement(l.Front(), 1)

  if result.Value != nthElement {
    t.Errorf("Result/Expected: %d/%d", result.Value, nthElement)
	}
}
