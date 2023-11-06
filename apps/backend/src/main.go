package main

import (
	"backend/api"
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on: http://localhost:9000")
	http.HandleFunc("/login", api.Login)
	http.ListenAndServe(":9000", nil)
}
