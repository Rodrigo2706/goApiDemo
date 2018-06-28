/**
* All Rights Reserved
* This software is proprietary information of Akurey
* Use is subject to license terms.
* Filename: validator.functions.go
*
* Author: rnavarro@akurey.com
* Description: Functions created for validation
* purposes
*/

package validators

import (
	"errors"
	"reflect"
)

/**
* Checks if a value is of type int
* @param  {interface{}}	value	Value to check
* @return (error)				If value is not of type int
*/
func isInt(value interface{}) error {
	var testInt int
	if reflect.TypeOf(value) != reflect.TypeOf(testInt) {
		return errors.New("value must be an int")
	}
	return nil
}