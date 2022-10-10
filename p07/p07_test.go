package p07

import (
	"99/util"
	"container/list"
	"testing"
)

func TestListFlatten(t *testing.T) {
	l := list.New()
	l.PushBack(1)
  l.PushBack(2)
  l.PushBack(3)

	ll := list.New()
  ll.PushBack(5)
  ll.PushBack(6)
  ll.PushBack(7)

	lll := list.New()
	lll.PushBack(0)
	lll.PushBack(l)
	lll.PushBack(4)
	lll.PushBack(ll)
	lll.PushBack(8)

	result := ListFlatten(lll)
	resultStr := util.ListToString(result.Front())
	expectedStr := "012345678"

  if resultStr != expectedStr {
    t.Errorf("Result/Expected: %s/%s", resultStr, expectedStr)
	}
}
