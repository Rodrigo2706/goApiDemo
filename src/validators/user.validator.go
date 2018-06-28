/**
* All Rights Reserved
* This software is proprietary information of Akurey
* Use is subject to license terms.
* Filename: user.validator.go
*
* Author: rnavarro@akurey.com
* Description: Handles all user validations
*/

package validators

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserValidator struct {
	Name       	string `json:"name"`
	Lastname 	string `json:"lastname"`
	Email		string `json:"email"`
}

type GetUsersValidator struct {
	Offset 	int `json:"offset"`
	Limit   int `json:"limit"`
}

/**
* Validates create user parameters
* @return {error}	If any
*/
func (user UserValidator) ValidateCreateUser() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required, validation.Length(1, 80)),
		validation.Field(&user.Lastname, validation.Required, validation.Length(1, 80)),
		validation.Field(&user.Email, validation.Required, is.Email),
	)
}

/**
* Validates get users parameters
* @return {error}	If any
*/
func (user GetUsersValidator) ValidateGetUsers() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Offset, validation.By(isInt)),
		validation.Field(&user.Limit, validation.By(isInt)),
	)
}

/**
* Validates get user by email parameters
* @return {error}	If any
*/
func (user UserValidator) ValidateGetUserByEmail() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required, is.Email),
	)
}

/**
* Validates update user parameters
* @return {error}	If any
*/
func (user UserValidator) ValidateUpdateUser() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Length(1, 80)),
		validation.Field(&user.Lastname, validation.Length(1, 80)),
		validation.Field(&user.Email, validation.Required, is.Email),
	)
}


