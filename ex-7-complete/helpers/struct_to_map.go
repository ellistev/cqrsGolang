package helpers

import (
	"reflect"
	"time"
)

func StructToMap(item interface{}) map[string]interface{} {

	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				if reflectValue.Field(i).Type() == reflect.TypeOf(time.Time{}) {
					res[tag] = reflectValue.Field(i).Interface().(time.Time).Format(time.RFC3339Nano)		// Error in .Interface()
				}else{
					res[tag] = StructToMap(field)
				}
			} else {
				res[tag] = field
			}
		}
	}
	return res
}