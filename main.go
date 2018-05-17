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

func melhorScrum(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Hello rodolpha1")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/danger", sayMyName)
	http.HandleFunc("/scrum", melhorScrum)
	http.ListenAndServe(":8000", nil)
}
