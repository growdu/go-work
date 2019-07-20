package net_test

import (
	"fmt"
	"net"
)

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
