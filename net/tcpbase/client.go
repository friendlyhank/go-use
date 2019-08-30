package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
)

func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if nil != err {
		fmt.Println(err)
		return
	}

	//2.发
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入你要发给服务器的内容:")
		line, _, _ := reader.ReadLine()
		conn.Write(line)
	}

	//1.收
	go ClientReadLoop(conn)
}

func ClientReadLoop(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, _ := conn.Read(buf)
		fmt.Println("【==>】", string(buf[0:n]))
	}
}

