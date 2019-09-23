package math

import (
	"math"
	"testing"
)

func TestMathMaxInt (t *testing.T){
	t.Logf("%v",math.MaxInt8)
	t.Logf("%v",math.MaxInt32)
	t.Logf("%v",math.MaxInt64)
}
