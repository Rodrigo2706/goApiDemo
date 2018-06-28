/**
* All Rights Reserved
* This software is proprietary information of Akurey
* Use is subject to license terms.
* Filename: empty.model.go
*
* Author: rnavarro@akurey.com
* Description: Declare the available properties
* of an Empty struct
*/

package models

import "github.com/nvellon/hal"

type EmptyStruct struct {
}

func (es EmptyStruct) GetMap() hal.Entry {
	return hal.Entry{
	}
}