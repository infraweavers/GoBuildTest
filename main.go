package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		time.Sleep(10 * time.Second)
		fmt.Fprintf(responseWriter, "Welcome to my website!")
	})
	http.ListenAndServe(":4080", nil)
}
