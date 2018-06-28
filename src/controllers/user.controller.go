/**
* All Rights Reserved
* This software is proprietary information of Akurey
* Use is subject to license terms.
* Filename: user.controller.go
*
* Author: rnavarro@akurey.com
* Description: Handles most of the logic of user
* creation, delete, update, get, etc.
*/

package controllers

import (
	"database/sql"
	"net/http"
	"encoding/json"
	"log"

	"github.com/nvellon/hal"
	"github.com/gorilla/mux"

	"validators"
	"helpers"
	"utils/common"
	"models"
)

type UserController struct {
	DB    *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{
		DB:    db,
	}

}

/**
* Handles the create user logic
* @param {http.ResponseWriter}	w	Used to send answer back
* @param {*http.Request}   		r  	Generated from API call
* @return
*/
func (userController *UserController) CreateUser(w http.ResponseWriter, r *http.Request){
	// Decode request body
	decoder := json.NewDecoder(r.Body)
	var userRequest validators.UserValidator

	err := decoder.Decode(&userRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate if parameters are correct
	err = userRequest.ValidateCreateUser()
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Send data to DB
	response, err := helpers.CreateUser(userController.DB, userRequest.Name, userRequest.Lastname, userRequest.Email)
	if err != nil || response != 200 {
		error := helpers.GetErrorByCode(response)
		http.Error(w, error, http.StatusBadRequest)
		return
	}

	// Send response back
	w.WriteHeader(http.StatusOK)
}

/**
* Handles get all users, works with pagination
* @param {http.ResponseWriter}	w	Used to send answer back
* @param {*http.Request}   		r  	Generated from API call
* @return
*/
func (userController *UserController) GetUsers(w http.ResponseWriter, r *http.Request){
	var err error
	var userRequest validators.GetUsersValidator

	path := r.URL.Path // The request path
	offset := 0 // Default values
	limit := 10 // Default values

	// Set values from URL Params
	err = common.SetParamToVar(r.URL.Query(), "offset", &offset)
	err = common.SetParamToVar(r.URL.Query(), "limit", &limit)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create user request to validate
	userRequest.Offset = offset
	userRequest.Limit = limit

	err = userRequest.ValidateGetUsers()
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get users from DB
	users, err := helpers.GetUsers(userController.DB, offset, limit)

	// Create hal response
	userResponse := hal.NewResource(models.EmptyStruct{}, path)
	for _, user := range users {
		halUser := hal.NewResource(user, path + "/" + user.Email)
		userResponse.Embed("users", halUser)
	}

	halJson, err := json.MarshalIndent(userResponse, "", "  ")

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.Write(halJson)
}

/**
* Handles getting a user from DB by Email
* @param {http.ResponseWriter}	w	Used to send answer back
* @param {*http.Request}   		r  	Generated from API call
* @return
*/
func (userController *UserController) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	var userRequest validators.UserValidator

	path := r.URL.Path // The request path

	// Create userRequest to validate
	params := mux.Vars(r)
	email := params["email"]
	userRequest.Email = email

	err := userRequest.ValidateGetUserByEmail()
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get user from DB
	user, err, response := helpers.GetUserByEmail(userController.DB, userRequest.Email)

	// Create HAL response
	userResponse := hal.NewResource(user, path)
	userResponse.AddNewLink("update_user", path)
	userResponse.AddNewLink("delete_user", path)

	halJson, err := json.MarshalIndent(userResponse, "", "  ")

	if err != nil || response != 200 || user == nil {
		error := helpers.GetErrorByCode(response)
		http.Error(w, error, http.StatusBadRequest)
		return
	}

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.Write(halJson)
}

/**
* Handles update a user by email
* @param {http.ResponseWriter}	w	Used to send answer back
* @param {*http.Request}   		r  	Generated from API call
* @return
*/
func (userController *UserController) UpdateUserByEmail(w http.ResponseWriter, r *http.Request) {
	var userRequest validators.UserValidator

	// Create userRequest to validate
	params := mux.Vars(r)
	email := params["email"]
	userRequest.Email = email

	// Decode request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate if parameters are correct
	err = userRequest.ValidateUpdateUser()
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update user in DB
	response, err := helpers.UpdateUserByEmail(userController.DB,
		userRequest.Email,
		userRequest.Name,
		userRequest.Lastname)

	if err != nil || response != 200 {
		error := helpers.GetErrorByCode(response)
		http.Error(w, error, http.StatusBadRequest)
		return
	}
	// Send response back
	w.WriteHeader(http.StatusOK)
}

/**
* Handles delete user from DB logic
* @param {http.ResponseWriter}	w	Used to send answer back
* @param {*http.Request}   		r  	Generated from API call
* @return
*/
func (userController *UserController) DeleteUserByEmail(w http.ResponseWriter, r *http.Request){
	var userRequest validators.UserValidator

	// Create userRequest to validate
	params := mux.Vars(r)
	email := params["email"]
	userRequest.Email = email

	// Validate if parameters are correct
	err := userRequest.ValidateUpdateUser()
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete user from DB
	response, err := helpers.DeleteUserByEmail(userController.DB,
		userRequest.Email)

	if err != nil || response != 200 {
		error := helpers.GetErrorByCode(response)
		http.Error(w, error, http.StatusBadRequest)
		return
	}
	// Send response back
	w.WriteHeader(http.StatusOK)
}