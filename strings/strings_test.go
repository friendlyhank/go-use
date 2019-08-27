package strings

import (
	"testing"

	"strings"
)

/**
Split和Join常用于数据库字段读写转化
*/

//Split -字符串分割,返回切片
func TestStringSplit(t *testing.T) {
	envstring := "abc,sss,sasa"
	args := strings.Split(envstring, ",")
	t.Log(args)
}

//Join- slice合并
func TestStringJoin(t *testing.T) {
	sSlice := []string{"aa", "bb", "cc"}
	joinStr := strings.Join(sSlice, ",")
	t.Log(joinStr)
}

//ToLower -字符串小写转化
func TestStringToLowwer(t *testing.T) {
	str := "AbdkT"
	tostr := strings.ToLower(str)
	t.Log(tostr)
}

//Contains -字符串包含于
func TestStringContain(t *testing.T) {
	stra := "abcd"
	strb := "b"
	ok := strings.Contains(stra, strb)
	t.Log(ok)

	strc := "be"
	ok = strings.Contains(stra, strc)
	t.Log(ok)
}
