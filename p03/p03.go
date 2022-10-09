package p03

import (
	"container/list"
)

func NthElement(e *list.Element, n int) *list.Element {
	if n == 0 {
		return e
	} else {
		return NthElement(e.Next(), n-1)
	}
}
