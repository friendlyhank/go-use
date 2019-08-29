package main

import (
	"log"
	"net"
)

/**
 *tcp基础面试中挺有用的
 */
func main(){

	//listen 建立套接字,绑定端口,监听端口
	l,err := net.Listen("tcp",":8088")
	if err != nil{
		log.Panic(err)
		return
	}

	for{
		conn,err := l.Accept()
		if err != nil{
			log.Panic(err)
			return
		}

		//maybe online func

		//握手过程 其实也是消息互通的过程
		HandleClientConn(conn)
	}
}

//HandleClientConn -
func HandleClientConn(conn net.Conn){
	//三大法宝 1.读取数据 2.处理 3.写数据
	go ReadLoop(conn)
	go HandleMessageLoop(conn)
	go WriteLoop(conn)
}

//读取数据
func ReadLoop(conn net.Conn){

	//maybe set timeout

}

func HandleMessageLoop(conn net.Conn){

}

//写数据
func WriteLoop(conn net.Conn){

}