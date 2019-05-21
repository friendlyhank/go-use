package md5

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/logs"
	"hash/crc32"
	"testing"
)

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//Md5加密
func TestMd5New(t *testing.T){
	strMd5 := GetMd5String("abc123")
	logs.Info("%v",strMd5)
}

//Crc32算法
func TestCrc32(t *testing.T){
	strCrc32 :=crc32.ChecksumIEEE([]byte("abc123"))
	logs.Info("%v",strCrc32)
}