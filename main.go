package main

import (
	"crud-golang-api/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/", handler.api)

	log.Println("localhost : 8000")
	http.ListenAndServe(":8000", nil)
}