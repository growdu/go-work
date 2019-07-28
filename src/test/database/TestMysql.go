package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/growdu/go-work/src/Utils"
	jsoniter2 "github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type ConnString struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func main() {
	file := `L:\gopath\src\github.com\growdu\go-work\src\conf\config.json`
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(dir)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	var connstring ConnString
	jsoniter := jsoniter2.ConfigCompatibleWithStandardLibrary
	err = jsoniter.Unmarshal(data, &connstring)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(connstring)
	//github.com/Go-SQL-Driver/MySQL
	con := connstring.User
	con += ":" + connstring.Password + "@tcp(" + connstring.Host + ":" + strconv.Itoa(connstring.Port) + ")"
	con += "/" + connstring.Dbname + "?charset=utf8"
	db, err := sql.Open("mysql", con)
	Utils.CheckError(err)
	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	Utils.CheckError(err)
	res, err := stmt.Exec("jack", "研发部门", "2012-12-09")
	Utils.CheckError(err)
	id, err := res.LastInsertId()
	Utils.CheckError(err)
	fmt.Println(id)
}
