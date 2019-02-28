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

	3.errorHandling有三个
	flag.ContinueOnError //错误但是会继续执行程序
	ExitOnError //错误然后终止程序
	PanicOnError //错误会抛出异常

	 */
	 //1.NewFlagSet new一个指令
	fs := flag.NewFlagSet("etcd",flag.ContinueOnError)

	//2.设置指令
	fs.String("config","","path to config file")
	fs.String("log-prefix","[nsqlookupd]","path to config file")
	fs.Bool("version",false,"path to version")
	err := fs.Parse(os.Args[1:])

	logs.Info("%v",err)

	//3.Help指令
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

	//4.输出某个具体的指令
	fmt.Println(fs.Arg(0))

	//5.指定命令输出
	/**
		Lookup查询具体某个指令,flag使用map去查询的
		里面接口的实现比较难理解，以后总结Hank
	 */
	if fs.Lookup("version").Value.(flag.Getter).Get().(bool){
		fmt.Println("nsqlookupd2.0")
		os.Exit(0)
	}

}

