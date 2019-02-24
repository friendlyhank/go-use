package os

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
	"strconv"
	"testing"
	"errors"
)

//打开文件
func TestOsOpen(*testing.T){
	file, err := os.Open("test.txt")
	if err != nil {
		logs.Info("%v",err)
	}
	logs.Info("%v", file)

	//open file: no such file or directory
}

//环境变量操作
func TestOsSetenv(*testing.T){
	//设置环境变量(并没有设置到环境变量中,应该是临时的)
	err := os.Setenv("TMPDOR","D:/logtest/")

	//获取环境变量
	show := func(key string){
		val,ok := os.LookupEnv(key)
		if !ok{
			fmt.Printf("%s not set\n", key)
		}else{
			fmt.Printf("%s=%s\n", key, val)
		}
	}

	show("GOPATH")
	show("USER")
	show("TMPDOR")

	//删除环境变量
	err = os.Unsetenv("TMPDOR")
	logs.Info("%v",err)

	//输入参数 例如:go run os.go 1 2
	for idx, args := range os.Args {
		fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
	}
}

//退出当前程序
func TestOsExit(*testing.T){
	os.Exit(0) //表示成功,退出
	os.Exit(1) //表示因为错误,退出
}

func TestOsStdErr(*testing.T){
	//错误输出到控制台
	err := errors.New("请求错误")
	fmt.Fprint(os.Stderr,err)
}


