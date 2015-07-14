package main

import (
	"io"
	"net/http"
  "os"
  "fmt"
)

import "github.com/DO-HW/dummy-go-lib"

func main() {
  http.HandleFunc("/", hello)

  myPort := os.Getenv("PORT")
  if (myPort == "") {
    myPort = "8088"
  }

  bind := fmt.Sprintf("%s:%s", os.Getenv("HOST"), myPort)
  fmt.Printf("listening on %s...", bind)
  err := http.ListenAndServe(bind, nil)
  if err != nil {
    panic(err)
  }
}


func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello "+ sayhello.HelloWorld() +"!")
}
