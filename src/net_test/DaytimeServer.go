package main

import (
	"github.com/growdu/go-work/src/Utils"
	"net"
	"time"
)

/*
 *daytime服务端，监听端口为1200
 *测试方式为 telnet localhost 1200
 * 每次会返回当前时间
 */
func main() {
	services := ":1200"
	//创建tcp地址
	tcpaddr, err := net.ResolveTCPAddr("tcp4", services)
	Utils.CheckError(err)
	//监听端口
	listener, err := net.ListenTCP("tcp", tcpaddr)
	Utils.CheckError(err)
	for {
		//死循环处理请求
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}
