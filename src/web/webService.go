package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
	"time"
)

type todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type todos []Todo
/*
web 测试，运行在10080端口，返回url地址
*/
func Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/home", home)
	router.HandleFunc("/todos", todoIndex)
	router.HandleFunc("/todos/{todoId}", todoShow)
	log.Fatal(http.ListenAndServe(":10080", router))
}


func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "你正在访问%q", html.EscapeString(r.URL.Path))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<body>This is home.</body>")
}

//func TodoIndex(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Todo Index!")
//}

func todoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func todoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}