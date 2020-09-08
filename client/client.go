package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// ip =  10.190.180.160
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client connect failed = ", err)
	}
	//fmt.Printf("connect 成功 = %v", conn)
	//从终端获取输入
	fmt.Println("连接服务器成功！！！")
	for {
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("reader err = ", err)
		}
		s := strings.Trim(str, "\r\n")
		_, err = conn.Write([]byte(str))
		if err != nil {
			fmt.Println("write err = ", err)
		}
		if s == "exit" {
			fmt.Println("客户端退出")
			break
		}
		//将输入法送给服务器端
		//fmt.Printf("客户端发送了%d字节的数据\n", n)
	}
}
