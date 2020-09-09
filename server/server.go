package main

import (
	"fmt"
	"net"
	"strings"
)

func process(conn net.Conn) {

	defer conn.Close() //关闭
	for {
		buf := make([]byte, 1024)
		//fmt.Printf("服务器在等待%v的输入\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if strings.Trim(string(buf[:n]), "\r\n") == "exit" {
			fmt.Println("客户端退出,等待新的客户连接")
			return
		}
		if n == 0 || err != nil {
			fmt.Println("出现不可预料的错误:", err)
			break
		}
		//在服务端输出
		//fmt.Print(string(buf[:n]))
		//重新发回给客户端
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("write err = ", err)
		}
	}
}

func main() {
	fmt.Println("服务器开始监听……")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen error = ", err)
		return
	}
	defer listen.Close() //延时关闭
	fmt.Println("等待客户来连接")
	//循环等待
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err = ", err)
		} else {
			fmt.Printf("连接成功，Accept() suc = %v 客户端 ip = %v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
	//fmt.Printf("listen suc= %v", listen)
}
