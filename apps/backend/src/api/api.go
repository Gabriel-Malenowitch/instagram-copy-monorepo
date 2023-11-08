package api

import (
	"backend/db"
	"backend/safety"
	"backend/utils"
	"encoding/json"
	"errors"
	"model"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUserById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	user := db.GetUserById(params["id"])
	byteArray, err := json.Marshal(user)
	utils.ThrowPanic(err)

	writer.Write(byteArray)
}

func GetPosts(writer http.ResponseWriter, request *http.Request) {
	posts := db.GetPosts()
	byteArray, err := json.Marshal(posts)
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

func Follow(writer http.ResponseWriter, request *http.Request) {
	var followPayload model.Followers
	err := json.NewDecoder(request.Body).Decode(&followPayload)
	utils.ThrowPanic(err)

	isValidPayload := safety.IsValidFollowPayload(followPayload)
	if !isValidPayload {
		utils.ThrowPanic(errors.New("invalid data"))
	}

	db.Follow(followPayload)
}

func Unfollow(writer http.ResponseWriter, request *http.Request) {
	var unfollowPayload model.Followers
	err := json.NewDecoder(request.Body).Decode(&unfollowPayload)
	utils.ThrowPanic(err)

	isValidPayload := safety.IsValidUnfollowPayload(unfollowPayload)
	if !isValidPayload {
		utils.ThrowPanic(errors.New("invalid data"))
	}

	db.Unfollow(unfollowPayload)
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
