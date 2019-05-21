package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"runtime"
	"testing"
	"time"
)

/**
Package runtime contains operations that interact with Go's runtime system, such as functions to control goroutines.

runtime会用到系统级别，然后还会去控制goroutines
 */

 //输出当前的程序架构
 //GOARCH支持386, amd64, arm, s390x
func TestRunTimeGoarch(t *testing.T){
	logs.Info("%v",runtime.GOARCH)
}

//返回系统信息 windows,darwin, freebsd, linux, and so on.
func TestRunTimeGoos(t *testing.T){
	logs.Info("%v",runtime.GOOS)
}

//获取go版本信息
func TestRunTimeGos(t *testing.T){
	logs.Info("%v",runtime.Version())
}

//获取CPU的数量
func TestRunTimeNumCPU(t *testing.T){
	logs.Info("%v",runtime.NumCPU())
}

//设置可执行CPU的数量
func TestRuntimeGomaxprocs(t *testing.T){
	logs.Info("%v",runtime.GOMAXPROCS(0))
}

func PrintA(){
	runtime.Gosched()
	fmt.Println("AAA")
}

func PrintB(){
	fmt.Println("BBB")
}

//释放当前gouroutine,执行下一个gouroutine
//函数的作用是暂停调用他的Goroutine的运行，调用他的Goroutine会被重新置于Gorunnable状态，并被放入调度器可运行G队列中
func TestRuntimeGosched(t *testing.T){
	go PrintA()
	go PrintB()
	time.Sleep(10 *time.Second)
}


