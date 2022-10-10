package util

import (
	"container/list"
	"testing"
)

func TestListToString(t *testing.T) {
	l := list.New()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	result := ListToString(l.Front())
	expected := "123"

	if result != expected {
    t.Errorf("Result/Expected: %s/%s", result, expected)
	}
}
