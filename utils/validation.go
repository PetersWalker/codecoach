//queryvars =  validator.validate(r, interface{})

package utils

import (
	"log"
)

// func Validate[T interface{ ~struct{ any } }](r http.Request, contraintMap map[any]func()) T {

// 	queryParams := r.URL.Query()

// 	t := T{}
// 	for k, _ := range queryParams {
// 		if contraintMap[k](k) { // e.g. notNull()

// 		}
// 		return k
// 	}

// 	return t
// }

func notNull(k any) {
	if k == nil {
		log.Panic()
	}

}
