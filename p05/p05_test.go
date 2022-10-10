package p05

import (
	"99/util"
	"container/list"
	"testing"
)

func TestListReverse(t *testing.T) {
	l := list.New()

	l.PushBack(1)
  l.PushBack(2)
	l.PushBack(3)

	result := ListReverse(l)
	expected := list.New()

	expected.PushBack(3)
  expected.PushBack(2)
	expected.PushBack(1)

	resultStr := util.ListToString(result.Front())
	expectedStr := util.ListToString(expected.Front())

  if resultStr != expectedStr {
    t.Errorf("Result/Expected: %s/%s", resultStr, expectedStr)
	}
}
