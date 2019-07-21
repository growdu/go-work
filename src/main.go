package main

import (
	"github.com/growdu/go-work/src/script"
	"github.com/growdu/go-work/src/web/gin"
)

//import "./web/gin"

func main() {
	//args:=[]string{"IP","192.168.1.124"}
	//net_test.ParseIP(args)
	//args[1]="0:0:0:0:0:0:0:1"
	//net_test.ParseIP(args)
	script.Start()
	gin.Start()

	//concurrentTest.Start()
	//mux.Run()
}
