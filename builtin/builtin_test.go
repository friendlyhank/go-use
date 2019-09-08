package builtin

import (
	"testing"
)

//
func TestCopy(t *testing.T){
	enport := []string{"2379","2380"}

	eps := make([]string, len(enport))

	copy(eps,enport)

	t.Logf("%v",eps)
}