/**
* All Rights Reserved
* This software is proprietary information of Akurey
* Use is subject to license terms.
* Filename: routes.go
*
* Author: rnavarro@akurey.com
* Description: Entry point for all routes
* accepted by the API
*/

package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
)

/**
* Creates all routes accepted by the API
* @param  {*mux.Router}	router	Router to add routes to
* @param  {*sql.DB}		db 		The db connection
* @return
*/
func CreateRoutes(router *mux.Router, db *sql.DB) {
	CreateUserRoutes(router, db)
}