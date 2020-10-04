package middleware

import (
	"fmt"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var valid *validator.Validate

func Validator(i interface{}) error {
	if valid == nil {
		valid = validator.New()
	}

	zh := zh.New()
	uni := ut.New(zh, zh)

	trans, _ := uni.GetTranslator("zh_Hans_CN")
	if err := zh_translations.RegisterDefaultTranslations(valid, trans); err != nil {
		return err
	}

	if err := valid.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			return fmt.Errorf(e.Translate(trans))
		}
	}
	return nil
}
