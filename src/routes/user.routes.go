/**
* All Rights Reserved
* This software is proprietary information of Akurey
* Use is subject to license terms.
* Filename: user.routes.go
*
* Author: rnavarro@akurey.com
* Description: Declares all User routes
*/

package routes

import (
	"database/sql"

	"github.com/gorilla/mux"

	"controllers"
)

const (
	USER_PATH = "/users"
)

/**
* Generates all user routes
* @param  {*mux.Router}	router	Router to add routes to
* @param  {*sql.DB}		db 		The db connection
* @return
*/
func CreateUserRoutes(router *mux.Router, db *sql.DB) {
	userController := controllers.NewUserController(db)

	// POST /users
	router.HandleFunc(USER_PATH, userController.CreateUser).Methods("POST")
	// GET /users/{email}
	router.HandleFunc(USER_PATH + "/{email}", userController.GetUserByEmail).Methods("GET")
	// GET /users
	router.HandleFunc(USER_PATH, userController.GetUsers).Methods("GET")
	// PATCH /users/{email}
	router.HandleFunc(USER_PATH + "/{email}", userController.UpdateUserByEmail).Methods("PATCH")
	// DELETE /users/{email}
	router.HandleFunc(USER_PATH + "/{email}", userController.DeleteUserByEmail).Methods("DELETE")
}
