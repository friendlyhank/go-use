package math

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestRandInit(t *testing.T){
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(2))
}

func TestMathMaxInt (t *testing.T){
	t.Logf("%v",math.MaxInt8)
	t.Logf("%v",math.MaxInt32)
	t.Logf("%v",math.MaxInt64)
}
