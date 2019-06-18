package mux

import (
	"log"
	"net/http"
)

func Run() {
	router := newRouter()

	log.Fatal(http.ListenAndServe(":10080", router))
}
