/**
* All Rights Reserved
* This software is proprietary information of Akurey
* Use is subject to license terms.
* Filename: user.model.go
*
* Author: rnavarro@akurey.com
* Description: Declare the available properties
* of a User struct
*/

package models

import "github.com/nvellon/hal"

type User struct {
	Name       	string `json:"name"`
	Lastname 	string `json:"lastname"`
	Email		string `json:"email"`
}

func (u User) GetMap() hal.Entry {
	return hal.Entry{
		"name":  u.Name,
		"lastname": u.Lastname,
		"email": u.Email,
	}
}