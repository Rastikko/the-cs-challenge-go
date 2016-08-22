package main

import (
	"log"
	"net/http"
)

func main() {
	AuthMain()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
