package _struct

import(
	"testing"
)

type Student struct{
}

func TestStructMapIsNil(t *testing.T){
	student := &Student{}

	t.Logf("%v",student == nil)
}
