package main

import (
	"fmt"
	jsoniter2 "github.com/json-iterator/go"
	"log"
)

type student struct {
	Id   int
	Name string
}

func main() {

	stu := student{
		Id:   1,
		Name: "jack",
	}
	var jsoniter = jsoniter2.ConfigCompatibleWithStandardLibrary
	data, err := jsoniter.Marshal(&stu)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	var jack student
	err = jsoniter.Unmarshal(data, &jack)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jack)
}
