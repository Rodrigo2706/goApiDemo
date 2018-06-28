/**
* All Rights Reserved
* This software is proprietary information of Akurey
* Use is subject to license terms.
* Filename: user.helper.go
*
* Author: rnavarro@akurey.com
* Description: Makes all user related DB calls
*/

package helpers

import (
	"database/sql"

	"models"
)

/**
* Creates a user un the DB
* @param  {*sql.DB}	db 			The db connection
* @param  {string}	pName 		The user name
* @param  {string}	pLastname	The user lastname
* @param  {string}	pEmail 		The user email
* @return (int, error)	Return code from DB, Error if any
*/
func CreateUser(db *sql.DB, pName, pLastname, pEmail string) (int, error) {
	const query = `CALL TUTSP_CreateUser(?, ?, ?)`
	var returnCode int
	rows, err := db.Query(query, pName, pLastname, pEmail)
	if rows.Next() {
		err = rows.Scan(&returnCode)
		if err != nil {
			return 0, err
		}
	}
	return returnCode, err
}

/**
* Gets a user from DB by its email
* @param  {*sql.DB}	db 			The db connection
* @param  {string}	pEmail 		The user email to get
* @return (*models.User, int, error)
	User data, Return code from DB if any, Error if any
*/
func GetUserByEmail(db *sql.DB, pEmail string) (*models.User, error, int){
	const query = `CALL TUTSP_GetUserByEmail(?)`
	var user models.User
	returnCode := 200

	rows, err := db.Query(query, pEmail)

	if err != nil {
		return nil, err, 0
	}
	if rows.Next() {
		err = rows.Scan(&returnCode)

		if err != nil && returnCode != 200 {
			return nil, err, returnCode
		} else {
			err = rows.Scan(&user.Name, &user.Lastname, &user.Email)

			if err != nil && returnCode != 200 {
				return nil, err, returnCode
			}
		}
	}
	return &user, err, returnCode
}

/**
* Creates a user un the DB
* @param  {*sql.DB}	db 			The db connection
* @param  {int}		pOffset		The page to get
* @param  {int}		pLimit		The number of records to get
* @return ([]models.User, error)	Array if user objects, Error if any
*/
func GetUsers(db *sql.DB,  pOffset, pLimit int) ([]models.User, error) {
	const query = `CALL TUTSP_GetUsers(?, ?)`

	users := make([]models.User, 0)

	rows, err := db.Query(query, pOffset, pLimit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Name, &user.Lastname, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

/**
* Updates a user in the DB
* @param  {*sql.DB}	db 			The db connection
* @param  {string}	pEmail 		The user email to update
* @param  {string}	pName 		The new user name
* @param  {string}	pLastname	The new user lastname
* @return (int, error)	Return code from DB, Error if any
*/
func UpdateUserByEmail(db *sql.DB, pEmail, pName, pLastname string) (int, error) {
	const query = `CALL TUTSP_UpdateUserByEmail(?, ?, ?)`
	var returnCode int
	rows, err := db.Query(query, pEmail, pName, pLastname)
	if rows.Next() {
		err = rows.Scan(&returnCode)

		if err != nil {
			return 0, err
		}
	}
	return returnCode, err
}

/**
* Deletes a user in the DB
* @param  {*sql.DB}	db 			The db connection
* @param  {string}	pEmail 		The user email to delete
* @return (int, error)	Return code from DB, Error if any
*/
func DeleteUserByEmail(db *sql.DB, pEmail string) (int, error) {
	const query = `CALL TUTSP_DeleteUserByEmail(?)`
	var returnCode int
	rows, err := db.Query(query, pEmail)
	if rows.Next() {
		err = rows.Scan(&returnCode)

		if err != nil {
			return 0, err
		}
	}
	return returnCode, err
}