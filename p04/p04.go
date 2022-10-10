package p04

import (
	"container/list"
)

func ListLength(l *list.List) int {
	return listLength(l.Front())
}

func listLength(e *list.Element) int {
	if e == nil {
		return 0
	} else {
		return 1 + listLength(e.Next())
	}
}
