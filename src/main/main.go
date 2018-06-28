/**
* All Rights Reserved
* This software is proprietary information of Akurey
* Use is subject to license terms.
* Filename: main.go
*
* Author: rnavarro@akurey.com
* Description: Entry point of the program
*/

package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"

	"utils/database"
	"routes"
)

func main() {
	// DB Connection
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	// Mux router
	router := mux.NewRouter()

	// Controller for routes
	routes.CreateRoutes(router, db)

	// Run server
	if err := http.ListenAndServe(":8000", router); err == nil {
		fmt.Println("Listeneing on port 8000")
	} else {
		log.Fatal(err)
	}
}