package main

import "fmt"

//声明接口
type test interface {
	String(input string)
}

//声明函数echo
type Echo func(input string)

//函数实现接口
func (e Echo) String(input string) {
	e(input)
}

func main() {
	//使用Echo()将echoInstance转为Echo类型
	echo := Echo(echoInstance)
	var test1 test = echo
	Print(echo)
	Print(test1)

}

func Print(test2 test) error {
	test2.String("test")
	return nil
}

func echoInstance(input string) {
	fmt.Println(input)
}
