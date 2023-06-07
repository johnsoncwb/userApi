package logDataNormalizer

import (
	"reflect"
)

func StructToMap(obj interface{}) map[string]interface{} {
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()

	// Make sure obj is a struct
	if objType.Kind() != reflect.Struct {
		panic("obj must be a struct")
	}

	data := make(map[string]interface{})
	for i := 0; i < objValue.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i)
		data[field.Name] = fieldValue.Interface()
	}

	return data
}
