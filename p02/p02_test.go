package p02

import (
	"container/list"
  "testing"
)

func TestLastElement(t *testing.T) {
	l := list.New()

	lastButOneElement := 99

	l.PushBack(1)
  l.PushBack(lastButOneElement)
	l.PushBack(3)

	result := LastButOneElement(l)

  if result.Value != lastButOneElement {
    t.Errorf("Result/Expected: %d/%d", result.Value, lastButOneElement)
	}
}
