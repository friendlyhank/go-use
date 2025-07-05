package main

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego/logs"
	"log"
	"os"
	"strings"
)

var (
	flagHelp = `
Member:
	--name 'default'
	--heartbeat-interval '100'
   `
)

/*
*
1.flag自带help命令
go run flag.go --help

2.也支持自定义NewFlagSet etcd
go run flag.go --etcd

3.errorHandling有三个
flag.ContinueOnError //错误但是会继续执行程序
ExitOnError //错误然后终止程序
PanicOnError //错误会抛出异常
*/
func NewFlagSet() {
	//1.NewFlagSet new一个指令
	fs := flag.NewFlagSet("etcd", flag.ContinueOnError)

	//2.设置指令
	fs.String("config", "", "path to config file")
	fs.String("log-prefix", "[nsqlookupd]", "path to config file")
	fs.Bool("version", false, "path to version")
	err := fs.Parse(os.Args[1:])

	log.Printf("%v", err)

	//3.Help指令
	if flag.ErrHelp == err {
		fmt.Println(flagHelp)
	}

	/**
	当只有输入--help或者添加的指令--etcd时,len(fs.Args()) == 0
	如果输入其他指令者不为零,此处可以检测无法识别的输入指令(除了注册的指令和系统自带指令)
	*/
	if len(fs.Args()) != 0 {
		logs.Info("%v", "无法识别指令")
	}

	//4.输出某个具体的指令
	fmt.Println(fs.Arg(0))

	//5.获取指定命令的输出
	fs.Lookup("version").Value.String()

	isVersion := fs.Lookup("version").Value.(flag.Getter).Get().(bool)
	fmt.Println(isVersion)
}

// 定义自定义类型
type StringSlice []string

// 实现 Value.String() 方法
func (s *StringSlice) String() string {
	return fmt.Sprintf("%v", *s)
}

// 实现 Value.Set() 方法
func (s *StringSlice) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

// FlagVar - 自定义复杂类型的命令解析方式
func FlagVar() {
	var list StringSlice
	flag.Var(&list, "list", "Comma-separated list of values")

	// 解析参数
	flag.Parse()

	// 输出结果
	fmt.Println("Parsed list:", list)
}
