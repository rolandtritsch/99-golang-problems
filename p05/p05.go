package p05

import (
	"container/list"
)

func ListReverse(l *list.List) *list.List {
	return listReverse(l.Front())
}

func listReverse(e *list.Element) *list.List {
	if e == nil {
		return list.New()
	} else {
		l := listReverse(e.Next())
		l.PushBack(e.Value)
		return l
	}
}
