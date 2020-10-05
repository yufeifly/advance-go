package main

import (
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	//w.Write([]byte("hello, world"))
	io.WriteString(w, "hello,world")
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("listen and serve: ", err)
	}
}
