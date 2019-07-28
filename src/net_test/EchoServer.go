package main

import (
	"fmt"
	"github.com/growdu/go-work/src/Utils"
	"net"
)

func main() {
	//SingleThread()
	multiThread()
}

func SingleThread() {
	service := ":1201"
	tcpaddr, err := net.ResolveTCPAddr("tcp4", service)
	Utils.CheckError(err)
	listener, err := net.ListenTCP("tcp", tcpaddr)
	Utils.CheckError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handleClient(conn)
		conn.Close()
	}

}

func multiThread() {
	service := "127.0.0.1:1202"
	tcpaddr, err := net.ResolveTCPAddr("tcp4", service)
	Utils.CheckError(err)
	listener, err := net.ListenTCP("tcp", tcpaddr)
	Utils.CheckError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClientMulti(conn)
		//conn.Close()
	}
}

func handleClientMulti(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		fmt.Println(string(buf[0:n]))
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func handleClient(conn net.Conn) {
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		fmt.Println(string(buf[0:n]))
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}
