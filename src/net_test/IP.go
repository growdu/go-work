package main

import (
	"fmt"
	"net"
)

func main(){
	args:=[]string{"IP","192.168.1.124"}

	ParseIP(args)
}

func ParseIP(args []string) {
	if len(args) != 2 {
		fmt.Println("Invalid error")
		return
	}
	name := args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address.")
		return
	} /*else{

		fmt.Println("The address is ",addr.String())
	}*/
	address,err:=net.ResolveIPAddr("ip",name)
	if err!=nil{
		return
	}
	print(address)
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()

	fmt.Println(
		"The address is ", addr.String(),
		"\nDefault mask length is ", bits,
		"\nLeading ones is ", ones,
		"\nMask hex is ", mask.String(),
		"\nNetwork is ", network.String())

}
