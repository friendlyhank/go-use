package net

import (
	"bufio"
	"net"
	"testing"
	"time"
)

//=============================================net socket=======================================

/**
*read write  listen.Accept不会保持阻塞的风格，所以要自己建立阻塞
*/

//TestNetListen -设置本地监听端口
func TestNetListen(t *testing.T){
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		t.Errorf("%v",err)
	}

	for{
		//接收每个链接
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

//获取IP地址
func TestInterfaceAddrs(t *testing.T){
	addrs,_ := net.InterfaceAddrs()

	for _,address := range addrs{
		ipnet, ok := address.(*net.IPNet)
		if !ok {
			continue
		}

		isUnspecified := ipnet.IP.IsUnspecified()
		if isUnspecified{
			t.Logf("%v","isUnspecified")
			t.Logf("地址%v",ipnet.IP.String())
		}

		//是否环回地址
		isLoopback :=ipnet.IP.IsLoopback()
		if isLoopback{
			t.Logf("%v","isLoopback")
			t.Logf("回环地址%v",ipnet.IP.String())
		}

		//多播地址
		isMulticast := ipnet.IP.IsMulticast()
		if isMulticast{
			t.Logf("%v","isMulticast")
			t.Logf("地址%v",ipnet.IP.String())
		}

		isInterfaceLocalMulticast := ipnet.IP.IsInterfaceLocalMulticast()
		if isInterfaceLocalMulticast{
			t.Logf("%v","isInterfaceLocalMulticast")
		}

		isLinkLocalMulticast := ipnet.IP.IsLinkLocalMulticast()
		if isLinkLocalMulticast{
			t.Logf("%v","isLinkLocalMulticast")
			t.Logf("地址%v",ipnet.IP.String())
		}

		isLinkLocalUnicast := ipnet.IP.IsLinkLocalUnicast()
		if isLinkLocalUnicast{
			t.Logf("%v","isLinkLocalUnicast")
			strIp :=ipnet.IP.String()
			t.Logf("地址%v",strIp)
		}
		isGlobalUnicast := ipnet.IP.IsGlobalUnicast()
		if isGlobalUnicast{
			t.Logf("%v","isGlobalUnicast")
			t.Logf("地址%v",ipnet.IP.String())
		}

	}
}

