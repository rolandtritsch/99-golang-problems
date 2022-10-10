package util

import (
	"container/list"
	"fmt"
)

func ListToString(e *list.Element) string {
	if e == nil {
		return ""
	} else {
	  i := e.Value.(int)
		return fmt.Sprintf("%d%s", i, ListToString(e.Next()))
	}
}
