package net

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net"
	"testing"
	"time"
)

//TestNetListen -设置本地监听端口
func TestNetListen(t *testing.T){
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for{
		conn,err := ln.Accept()
		if err != nil{
			//handle error
		}
		logs.Info("%v",conn)

		//输出地址listen.Addr
		logs.Info("%v",ln.Addr())
		//
		logs.Info("%v",conn.RemoteAddr())
	}
}

//TestNetDial -Dial拨号功能连接服务器
func TestNetDial(t *testing.T){
	conn,err := net.Dial("tcp","127.0.0.1:8080")
	fmt.Fprintf(conn,"GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	logs.Info("%v",err)
	logs.Info("%v",status)
}

//TestNetDialTimeOut -Dial拨号功能连接服务器
func TestNetDialTimeOut(t *testing.T){
	//1秒的超时呼叫
	conn,err := net.DialTimeout("tcp","127.0.0.1:8080",time.Second)
	fmt.Fprintf(conn,"GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	logs.Info("%v",err)
	logs.Info("%v",status)
}


