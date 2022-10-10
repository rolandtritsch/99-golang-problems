package p06

import (
	"container/list"
	"testing"
)

func TestListIsPalindrome(t *testing.T) {
	l := list.New()

	l.PushBack(3)
  l.PushBack(99)
  l.PushBack(99)
	l.PushBack(3)

	result := ListIsPalindrome(l)

  if !result {
    t.Errorf("Result/Expected: %t/%t", result, true)
	}
}
