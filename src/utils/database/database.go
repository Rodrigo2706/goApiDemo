/**
* All Rights Reserved
* This software is proprietary information of Akurey
* Use is subject to license terms.
* Filename: database.go
*
* Author: rnavarro@akurey.com
* Description: Handles mysql DB connection
*/

package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const(
	HOST = "localhost"
	PORT = 3305
	USER = "root"
	PASSWORD = "root"
	DBNAME = "GoApiTutorial"
)

/**
* Connects to DB
* @return (*sql.DB, error)	Pointer to DB connection, error if any
*/
func Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		USER, PASSWORD, HOST, PORT, DBNAME)
	return sql.Open("mysql", connStr)
}