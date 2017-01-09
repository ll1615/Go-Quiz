package lib

import (
	"fmt"
	"time"
)

// Trace a function entering moment and elapsed time.
// It should be called in the very beginning in a function,
// like this: defer lib.Trace()()
func Trace() func() {
	start := time.Now()
	fmt.Println(start)
	return func() { fmt.Println(time.Since(start)) }
}
