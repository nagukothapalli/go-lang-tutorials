package main

import (
	"fmt"

	"net/http"
)

type Hello struct{}

func (hello Hello) serveHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> hello welcome to Server ")
}

func main() {
	var hello Hello
	err := http.ListenAndServe("127.0.0.1:8181", hello)
	checkError(err)
}

func checkError(isTheError error) {
	if isTheError != nil {
		panic(isTheError)
	}
}
