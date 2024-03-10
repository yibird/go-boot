package utils

import (
	"fmt"
	"reflect"
)

// SetFieldValue 通过反射设置对象的指定属性值
func SetFieldValue(obj interface{}, fieldName string, value interface{}) error {
	// 获取对象的反射值
	objValue := reflect.ValueOf(obj).Elem()

	// 获取字段
	field := objValue.FieldByName(fieldName)

	// 检查字段是否存在
	//if !field.IsValid() {
	//	return fmt.Errorf("Field %s not found in struct", fieldName)
	//}

	// 检查字段是否可以被修改
	if !field.CanSet() {
		return fmt.Errorf("Cannot set value for unexported field %s", fieldName)
	}

	// 将值转换为字段的类型
	newValue := reflect.ValueOf(value)
	if newValue.Type() != field.Type() {
		return fmt.Errorf("Type mismatch for field %s", fieldName)
	}

	// 设置字段的值
	field.Set(newValue)

	return nil
}
