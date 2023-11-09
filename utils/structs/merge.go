package structs

import (
	"reflect"
)

// 合并两个结构体的同名属性值
func Merge(target interface{}, source interface{}) {
	targetType := reflect.TypeOf(target)
	sourceType := reflect.TypeOf(source)

	if targetType.Kind() == reflect.Ptr {
		targetType = targetType.Elem()
	}
	if sourceType.Kind() == reflect.Ptr {
		sourceType = sourceType.Elem()
	}
	for i := 0; i < targetType.NumField(); i++ {
		targetField := targetType.Field(i)
		sourceField, ok := sourceType.FieldByName(targetField.Name)
		if ok && sourceField.Name == targetField.Name {
			sourceValue := reflect.ValueOf(source)
			targetValue := reflect.ValueOf(target).Elem()
			if sourceValue.Kind() == reflect.Ptr {
				sourceValue = sourceValue.Elem()
			}
			if targetValue.Kind() == reflect.Ptr {
				targetValue = targetValue.Elem()
			}
			value := sourceValue.FieldByName(targetField.Name)
			targetValue.FieldByName(sourceField.Name).Set(value)
		}
	}
}
