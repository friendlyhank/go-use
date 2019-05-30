package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main(){
	//OsFlag()
	OsSignal()
}

//输入参数 例如:go run os.go 1 2
func OsFlag(){
	for idx, args := range os.Args {
		fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
	}
}

//signal信号处理
func OsSignal(){
	signRecv := make(chan os.Signal,1)

	//NOT []*os.Signal,取值
	//这样就可以自行处理SIGINT,SIGQUIT信号
	sigs := []os.Signal{syscall.SIGINT,syscall.SIGQUIT}
	signal.Notify(signRecv,sigs...)

	//试图用for语句接受信号值，只要signRecv中存在元素值,for就会按顺序接收并迭代变量sig,否则for就会阻塞等待新元素
	//如果signRecv通道关闭,for语句会立刻退出,所以不用担心程序会循环往复
	for s :=range signRecv{
		fmt.Printf("Received a signal: %s\n",s)
	}
}


