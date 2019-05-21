package io

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"os"
	"testing"
)

//读取指定大小内容
func TestIoReadFull(t *testing.T){
	r,_ := os.Open("../test.txt")
	buf := make([]byte,1)
	io.ReadFull(r,buf)
	logs.Info("%v",string(buf))
}

//TestIoGetwd -获取当前目录,类似linux的pwd
//runtime.GOOS会根据系统判断
//hank-sure
func TestIoGetwd(t *testing.T){
	pwd,_ := os.Getwd()
	fmt.Println(pwd)
}


