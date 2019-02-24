package main

import (
	"github.com/astaxie/beego/logs"
	"runtime"
	"testing"
)

/**
Package runtime contains operations that interact with Go's runtime system, such as functions to control goroutines.

runtime会用到系统级别，然后还会去控制goroutines
 */

 //输出当前的程序架构
 //GOARCH支持386, amd64, arm, s390x
func TestRunTimeGoarch(*testing.T){
	logs.Info("%v",runtime.GOARCH)
}


