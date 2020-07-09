package time

import (
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	// sleep
	time.Sleep(10 * time.Millisecond)
}
