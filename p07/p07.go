package p07

import (
	"container/list"
)

func ListFlatten(l *list.List) *list.List {
	return listFlatten(l.Front(), list.New())
}

func listFlatten(e *list.Element, l *list.List) *list.List {
	if e == nil {
		return l
	} else {
		switch e.Value.(type) {
		case *list.List:
			ll := e.Value.(*list.List)
			return listFlatten(e.Next(), listFlatten(ll.Front(), l))

		default:
			l.PushBack(e.Value)
			return listFlatten(e.Next(), l)
		}
	}
}
