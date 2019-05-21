package time

import (
	"testing"
	"time"
)

//相当于time.Now().Sub(t)
func TestTimeSince(t *testing.T) {
	elapsed := time.Since(time.Now())
	t.Logf("%v",elapsed)
}
