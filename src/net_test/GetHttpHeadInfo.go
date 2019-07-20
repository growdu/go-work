package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	start()
}

func start() {
	tcpType := "tcp4"
	service := "www.baidu.com:80"
	tcpAddr, err := net.ResolveTCPAddr(tcpType, service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
