package main

import (
	"backend/api"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func internalProxy(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	callback := func(writer http.ResponseWriter, request *http.Request) {
		// Do something...
		handler(writer, request)
	}

	return callback
}

func main() {
	log.Println("Server started on: http://localhost:9000")
	router := mux.NewRouter()

	http.HandleFunc("/login", internalProxy(api.Login))
	router.HandleFunc("/users/{id:.+}", SingleUser).Methods("GET", "PUT")
	router.HandleFunc("/configurations/{id:.+}", SingleConfiguration).Methods("GET", "PUT")
	router.HandleFunc("/users", internalProxy(api.InsertUser)).Methods("POST")

	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func SingleUser(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		internalProxy(api.GetUserById)(writer, request)
	case "PUT":
		internalProxy(api.UpdateUser)(writer, request)
	}
}

func SingleConfiguration(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		internalProxy(api.GetUserConfigurationById)(writer, request)
	case "PUT":
		internalProxy(api.UpdateUserConfiguration)(writer, request)
	}
}
