/**
* All Rights Reserved
* This software is proprietary information of Akurey
* Use is subject to license terms.
* Filename: error.helper.go
*
* Author: rnavarro@akurey.com
* Description: Maps DB errors to a readable string
*/

package helpers

const (
	EmailAlreadyExists	= 10000
	UserNotFound = 10001
)

// Maps error codes with readable message
var statusText = map[int]string{
	EmailAlreadyExists: "User email already exists",
	UserNotFound: "User email does not exist",
}

/**
* Changes an error code to a readable string
* @param  {string}	pCode The error code
* @return {string}		  The readable message
*/
func GetErrorByCode(pCode int) string{
	return statusText[pCode]
}