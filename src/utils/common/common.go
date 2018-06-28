/**
* All Rights Reserved
* This software is proprietary information of Akurey
* Use is subject to license terms.
* Filename: common.go
*
* Author: rnavarro@akurey.com
* Description: Common functions used in several places
*/

package common

import (
	"strconv"
	"net/url"
	"errors"
)

/**
* Takes URL params and assigns them to a variable
* @param  {url.Values}	pParams		Params array
* @param  {string}		pValueToGet	Parameter name to get
* @param  {*int}		pVariable	Variable to change
* @return {error}					If any
*/
func SetParamToVar(pParams url.Values, pValueToGet string, pVariable *int) error{
	param, ok := pParams[pValueToGet]
	if ok {
		result, err := strconv.Atoi(param[0])
		if err == nil {
			*pVariable = result
			return nil
		}
		return errors.New("could not assign value to int")
	}
	return nil
}