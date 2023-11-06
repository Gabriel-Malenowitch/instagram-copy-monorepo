package api

import (
	"fmt"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	body := request.Body
	fmt.Println("Payload:", body)
}
