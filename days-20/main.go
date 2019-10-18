package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func main() {
	var h Hello
	err := http.ListenAndServe("127.0.0.1:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}
