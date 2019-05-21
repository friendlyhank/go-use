package bytes

import (
	"bytes"
	"testing"
)

//bytes相等
func TestBytesEqual(t *testing.T){
	resp := []byte("123456")

	if bytes.Equal(resp,[]byte("123456")){
		t.Logf("%v","Equal")
	}else{
		t.Logf("%v","Not Equal")
	}
}