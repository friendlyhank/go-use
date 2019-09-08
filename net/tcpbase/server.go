package main

import (
	"fmt"
	"log"
	"net"
)


type Server struct{
	conn net.Conn

	//接收的消息
	receverMsg chan string
	writeMsg chan string

	exit bool
}

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

		server := &Server{
			conn:conn,
			receverMsg:make(chan string,0),
			writeMsg:make(chan string,0),
		}

		fmt.Println("server conning")

		//maybe online func

		//握手过程 其实也是消息互通的过程
		HandleClientConn(server)
	}
}

//HandleClientConn -
func HandleClientConn(server *Server){
	//三大法宝 1.读取数据 2.处理 3.写数据
	go ReadLoop(server)
	go HandleMessageLoop(server)
	go WriteLoop(server)
}

//读取数据
func ReadLoop(server *Server){

	//maybe set timeout
	for{
		buf := make([]byte,1024)
		n,err :=server.conn.Read(buf)
		if err != nil{
			server.conn.Close()
			return
		}

		msg := string(buf[:n])
		server.receverMsg <-  msg
	}
}

//处理数据
func HandleMessageLoop(server *Server){

	for{
		select {
			case receverMsg := <- server.receverMsg:
				//deal receverMsg

				resMsg :=fmt.Sprintf("接收到数据,deal %v",receverMsg)
				server.writeMsg <- resMsg
		}
	}
}

//写数据
func WriteLoop(server *Server){
	//maybe set timeout
	for{
			select{
				case writeMsg := <-  server.writeMsg:
					wb := []byte(writeMsg)
					_,err :=server.conn.Write(wb)
					if err != nil{
						server.conn.Close()
					}
			}
	}
}