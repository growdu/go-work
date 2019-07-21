package main

import (
	"encoding/asn1"
	"fmt"
	"github.com/growdu/go-work/src/Utils"
)

func main() {
	data, err := asn1.Marshal(13)
	fmt.Println(data)
	Utils.CheckError(err)
	var n int
	_, err1 := asn1.Unmarshal(data, &n)
	Utils.CheckError(err1)
	fmt.Println(n)
}
