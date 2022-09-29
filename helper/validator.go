package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var uni *ut.UniversalTranslator

func Validate(i interface{}) error {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ := getTranslator("en")

	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}
		return name
	})

	en_translations.RegisterDefaultTranslations(v, trans)

	if err := v.Struct(i); err != nil {
		return NewErr(http.StatusBadRequest, ErrorInvalidParameter, getValidatorMessage(err))
	}

	return nil
}

func getTranslator(lang string) (ut.Translator, error) {
	trans, _ := uni.GetTranslator(lang)
	return trans, nil
}

func getValidatorMessage(err error) string {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return ""
	}

	errs := make([]string, 0)
	for _, err := range err.(validator.ValidationErrors) {
		tmp := []string{fmt.Sprintf("validation failed on field %s with precondition '%s'", err.Field(), err.ActualTag())}

		if err.Param() != "" {
			if err.ActualTag() == "oneof" {
				tmp = append(tmp, fmt.Sprintf("want %s", strings.Replace(err.Param(), `' '`, `' or '`, -1)))
			} else {
				tmp = append(tmp, fmt.Sprintf("want %s", err.Param()))
			}
		}

		if err.Value() != nil && err.Value() != "" {
			tmp = append(tmp, fmt.Sprintf("but got %v", err.Value()))
		}

		errs = append(errs, strings.Join(tmp, " "))
	}

	b, _ := json.Marshal(errs)
	return string(b)
}
