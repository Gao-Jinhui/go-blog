package util

import "reflect"

func ConvertToMap(request interface{}) map[string]interface{} {
	typeOfRequest := reflect.TypeOf(request)
	valueOfRequest := reflect.ValueOf(request)
	filter := make(map[string]interface{})
	for pos := 0; pos < typeOfRequest.NumField(); pos++ {
		if !valueOfRequest.Field(pos).IsZero() {
			filter[typeOfRequest.Field(pos).Name] = valueOfRequest.Field(pos).Interface()
		}
	}
	return filter
}
