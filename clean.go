package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)



func main() {


	var msges string
	flag.StringVar(&msges, "u", "", "信息")
	flag.Parse()
	Sendmsg(msges)
}


func Sendmsg(msg string){

	//通过tcp 协议链接 本机 8080端口
	con, err := net.Dial("tcp", "192.168.1.14:8080")
	//如果出现错误 说明链接失败
	if err != nil {
		fmt.Println("连接服务器端失败")
		fmt.Println(err.Error())
		os.Exit(0)
	}

	//记得关闭
	defer con.Close()
	//开始向服务器端发送 hello
	num, write_err := con.Write([]byte(msg))
	//如果写入有问题 输出对应的错误信息
	if write_err != nil {
		fmt.Println(write_err.Error())
	}
	//如果没有问题。显示对应的写入长度
	fmt.Println(num)
}
