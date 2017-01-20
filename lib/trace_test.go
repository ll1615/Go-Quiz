package lib

import (
	"testing"
	"time"
)

func TestTrace(t *testing.T) {
	defer Trace()()
	time.Sleep(1 * time.Second)
}
