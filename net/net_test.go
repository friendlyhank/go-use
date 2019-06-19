package net

import (
	"bufio"
	"net"
	"testing"
	"time"
)

//TestNetListen -设置本地监听端口
func TestNetListen(t *testing.T){
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		t.Errorf("%v",err)
		}
	for{
		conn,err := ln.Accept()
		if err != nil{
			//handle error
			t.Errorf("%v",err)
			return
		}
		t.Logf("%v",conn)

		//输出地址listen.Addr
		t.Logf("%v",ln.Addr())

		//
		t.Logf("%v",conn.RemoteAddr())
	}
}

//TestNetDial -Dial拨号功能连接服务器
func TestNetDial(t *testing.T){
	conn,err := net.Dial("tcp","127.0.0.1:8080")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil{
		t.Errorf("%v",err)
		return
	}
	t.Logf("%v",status)
}

//TestNetDialTimeOut -Dial拨号功能连接服务器
func TestNetDialTimeOut(t *testing.T){
	//1秒的超时呼叫
	conn,err := net.DialTimeout("tcp","127.0.0.1:8080",time.Second)
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil{
		t.Errorf("%v",err)
		return
	}
	t.Logf("%v",status)
}


