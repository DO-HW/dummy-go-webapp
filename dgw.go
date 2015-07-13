package main

import (
	"io"
	"net/http"
)

import "github.com/DO-HW/dummy-go-lib"

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello "+ sayhello.HelloWorld() +"!")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
