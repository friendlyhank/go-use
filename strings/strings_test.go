package strings

import (
	"fmt"
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

//字符串小写转化
func TestStringToLowwer(t *testing.T){
	str := "AbdkT"
	tostr := strings.ToLower(str)
	fmt.Println(tostr)
}