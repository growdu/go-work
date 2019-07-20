package main

import "fmt"

func main() {
	//var test[]string
	var test = make([]string, 10)
	fmt.Printf("%p\n", test)
	test = append(test, "test1")
	//p:=&test
	fmt.Printf("%p\n", test)
	test = append(test, "test1")
	fmt.Printf("%p\n", test)
}
