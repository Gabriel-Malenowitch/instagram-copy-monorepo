package api

import (
	"backend/db"
	"backend/safety"
	"backend/utils"
	"encoding/json"
	"errors"
	"fmt"
	"model"
	"net/http"

	"github.com/gorilla/mux"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	body := request.Body
	fmt.Println("Payload:", body)
}

func GetUserById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	user := db.GetUserById(params["id"])
	byteArray, err := json.Marshal(user)
	utils.ThrowPanic(err)

	writer.Write(byteArray)
}

func GetUserConfigurationById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	configuration := db.GetUserConfigurationById(params["id"])
	byteArray, err := json.Marshal(configuration)
	utils.ThrowPanic(err)

	writer.Write(byteArray)
}

func InsertUser(writer http.ResponseWriter, request *http.Request) {
	var user model.UserInsertPayload
	err := json.NewDecoder(request.Body).Decode(&user)
	utils.ThrowPanic(err)

	isValidPayload := safety.IsValidInsertUserPayload(user)
	if !isValidPayload {
		utils.ThrowPanic(errors.New("invalid data"))
	}

	db.InsertUser(user)
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userId := params["id"]
	var user model.UserUpdatePayload
	err := json.NewDecoder(request.Body).Decode(&user)
	utils.ThrowPanic(err)

	isValidPayload := safety.IsValidUpdateUserPayload(user)
	if !isValidPayload {
		utils.ThrowPanic(errors.New("invalid data"))
	}

	db.UpdateUser(user, userId)
}

func UpdateUserConfiguration(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userId := params["id"]
	var configuration model.UserConfigurationPayload
	err := json.NewDecoder(request.Body).Decode(&configuration)
	utils.ThrowPanic(err)

	isValidPayload := safety.IsValidUpdateUserConfigurationPayload(configuration)
	if !isValidPayload {
		utils.ThrowPanic(errors.New("invalid data"))
	}

	db.UpdateUserConfiguration(configuration, userId)
}
