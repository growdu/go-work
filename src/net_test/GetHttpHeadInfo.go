package main

import (
	"fmt"
	"github.com/growdu/go-work/src/Utils"
	"io/ioutil"
	"net"
	)

func main() {
	start()
}

func start() {
	tcpType := "tcp4"
	service := "www.baidu.com:80"
	tcpAddr, err := net.ResolveTCPAddr(tcpType, service)
	Utils.CheckError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	Utils.CheckError(err)
	result, err := ioutil.ReadAll(conn)
	Utils.CheckError(err)
	fmt.Println(string(result))
}


