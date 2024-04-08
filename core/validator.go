package core

import (
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translator "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"strings"
)

var internalErrors = map[string]string{
	"required": "%s不能为空",
	"min":      "字段长度不能少于 %d 个字符",
	"max":      "字段长度不能超过 %d 个字符",
}

type Validator struct {
	validate *validator.Validate
	trans    ut.Translator
}

// NewValidator 创建一个自定义错误消息的验证器
func NewValidator() (*Validator, error) {
	validate := validator.New()
	// 创建翻译器实例
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	// 注册翻译器
	if err := translator.RegisterDefaultTranslations(validate, trans); err != nil {
		return nil, err
	}
	v := &Validator{
		validate: validate,
		trans:    trans,
	}
	v.RegisterCustomErrors(internalErrors)
	return v, nil
}

// Validate 执行结构体的验证，并返回验证错误
func (cv *Validator) Validate(i interface{}) error {
	if err := cv.validate.Struct(i); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// 获取标签上的自定义错误消息
			tag := err.Tag()
			fieldName := err.Field()
			// 获取标签对应错误信息
			errorMsg := getFieldError(i, fieldName, tag)
			if len(errorMsg) > 0 {
				return errors.New(errorMsg)
			}
			translate := err.Translate(cv.trans)
			if len(translate) > 0 {
				return errors.New(translate)
			}
			return errors.New(err.Error())
		}
	}
	return nil
}

// RegisterCustomErrors 注册自定义的错误消息映射
func (cv *Validator) RegisterCustomErrors(errors map[string]string) {
	for tag, message := range errors {
		err := cv.validate.RegisterTranslation(tag, cv.trans, func(ut ut.Translator) error {
			return ut.Add(tag, message, true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(fe.Tag(), fe.Field())
			return fmt.Sprintf(t, strutil.CamelCase(fe.Field()))
		})
		if err != nil {
			return
		}
	}
}

// getFieldError 获取字段上的自定义错误消息
func getFieldError(v interface{}, fieldName, tag string) string {
	// 使用反射获取结构体字段的标签值
	field, _ := reflect.TypeOf(v).FieldByName(fieldName)
	tagValue := field.Tag.Get("error")
	// 解析标签值,获取自定义错误消息
	return parseTag(tagValue, tag)
}

// parseTag 解析标签值，获取对应标签的错误消息
func parseTag(tagValue, tag string) string {
	parts := strings.FieldsFunc(tagValue, func(r rune) bool {
		return r == '|' || r == '='
	})
	for i := range parts {
		if parts[i] == tag && i < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}
