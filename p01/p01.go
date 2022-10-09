package p01

import (
	"container/list"
)

func LastElement(l *list.List) *list.Element {
	if l.Len() == 0 {
		return nil
	} else {
		return l.Back()
	}
}
