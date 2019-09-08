package net

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
)

/**
 *http代理
 *使用:浏览器开启代理，127.0.0.1 port:8081
 *场景：运行外国外VPS上,设置代理即可翻墙
 */
func httpPrivoxy(){

	//设置监听 listen
	l, err := net.Listen("tcp", ":8081")

	if err != nil{
		log.Panic(err)
	}


	for{
		//Accept
		client,err := l.Accept()

		if err != nil{
			log.Panic(err)
		}

		//go Handle
		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn){
	if client == nil{
		return
	}
	defer client.Close()

	//读取请求
	var b [1024]byte
	n,err := client.Read(b[:])
	if err != nil{
		fmt.Println("|字节数|%v|%v",n,string(b[:]))
		log.Println(err)
		return
	}



	var method,host,address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)

	hostPortURL,err :=url.Parse(host)
	if err != nil{
		fmt.Println("|字节数|%v|%v",n,string(b[:]))
		fmt.Println("host:",host)
		log.Println(err)
		return
	}

	//https访问
	if hostPortURL.Opaque == "443"{
		address = hostPortURL.Scheme + ":443"
	}else{
		if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}

	fmt.Println("address:",address)

	//获得了请求的host和post,就开始拨号
	server,err := net.Dial("tcp",address)
	if err != nil{
		log.Fatal(err)
		return
	}
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(b[:n])
	}

	//进行转发 如何进行转发是重点
	go io.Copy(server,client)
	io.Copy(client, server)
}