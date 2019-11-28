package encoding

import (
	"encoding/base64"
	"testing"
)

func TestEncodeBase64(t *testing.T){
	codeStr := base64.StdEncoding.EncodeToString([]byte("你好"))
	t.Logf("%v",codeStr)//5L2g5aW9
}

func TestDecodeBase64(t *testing.T){
	decideByte,_ := base64.StdEncoding.DecodeString("5L2g5aW9")
	t.Logf("%v",string(decideByte))
}
