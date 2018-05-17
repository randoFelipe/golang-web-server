package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "The cake is a lie!")
}

func sayMyName(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Hello Heisenberg")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/danger", sayMyName)
	http.ListenAndServe(":8000", nil)
}
