package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pong)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func pong(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}
