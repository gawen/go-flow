package flow

import (
	"fmt"
)

func panicf(msg string, args ...interface{}) {
	panic(fmt.Sprintf(msg, args...))
}
