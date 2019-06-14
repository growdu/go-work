package web

import (
	"log"
	"net/http"
)

func Run(){
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":10080", router))
}
