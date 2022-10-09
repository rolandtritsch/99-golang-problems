package p02

import (
	"container/list"
)

func LastButOneElement(l *list.List) *list.Element {
	if l.Len() == 0 {
		return nil
	} else {
		return l.Back().Prev()
	}
}
