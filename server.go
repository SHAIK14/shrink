package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello go\n")
}

func PersonCreate(w http.ResponseWriter, r *http.Request) {
	var p Person
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Person:%+v", p)

}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/person", PersonCreate)
	http.ListenAndServe(":8080", nil)
}
