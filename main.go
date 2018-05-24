package main

import (
	"io"
	"net/http"
	"os"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "mudou task definition")
}

func test_service_discovery(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://google.com")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	
	bs := make ([]byte, 9999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
	
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
	http.HandleFunc("/dnsstatus", test_service_discovery)
	http.ListenAndServe(":8000", nil)
}
