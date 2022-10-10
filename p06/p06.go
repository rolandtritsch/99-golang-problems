package p06

import (
	"99/p05"
	"99/util"
	"container/list"
)

func ListIsPalindrome(l *list.List) bool {
	return util.ListToString(l.Front()) == util.ListToString(p05.ListReverse(l).Front())
}
