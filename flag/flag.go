package main

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
)

func main(){
	FlagSetParse()
}

var (
	flagHelp = `
Member:
	--name 'default'
	--heartbeat-interval '100'
   `
)

func FlagSetParse(){
	/**
	1.flag自带help命令
	go run flag.go --help

	2.也支持自定义NewFlagSet etcd
	go run flag.go --etcd

	 */
	fs := flag.NewFlagSet("etcd",flag.ContinueOnError)
	err := fs.Parse(os.Args[1:])

	logs.Info("%v",err)

	//Help指令
	if flag.ErrHelp == err {
		fmt.Println(flagHelp)
	}

	/**
		当只有输入--help或者添加的指令--etcd时,len(fs.Args()) == 0
		如果输入其他指令者不为零,此处可以检测无法识别的输入指令(除了注册的指令和系统自带指令)
	 */
	if(len(fs.Args()) != 0){
		logs.Info("%v","无法识别指令")
	}

	//输出某个具体的指令
	fmt.Println(fs.Arg(0))
}

