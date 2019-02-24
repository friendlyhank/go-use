package strings

import (
	"github.com/astaxie/beego/logs"
	"strings"
	"testing"
)

//字符串分割,返回切片
func TestStringSplit(t *testing.T){
	envstring := "abc,sss,sasa"
	args := strings.Split(envstring,",")
	logs.Info("%v",args)
}
