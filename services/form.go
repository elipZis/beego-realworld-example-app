package services

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"net/url"
)

type FormService struct {
}

type FormError struct {
	Key     string
	Message string
}

func (this *FormService) ParseAndValidate(input url.Values, obj interface{}) (interface{}, map[string]FormError) {
	valid := validation.Validation{}
	errors := make(map[string]FormError)
	if err := beego.ParseForm(input, obj); err != nil {
		errors["Error"] = FormError{
			err.Error(),
			err.Error(),
		}
	}
	b, err := valid.Valid(obj)
	if err != nil {
		errors["Error"] = FormError{
			err.Error(),
			err.Error(),
		}
	}
	if !b {
		for _, err := range valid.Errors {
			errors[err.Field] = FormError{
				err.Key,
				err.Message,
			}
		}
	}

	if len(errors) > 0 {
		return obj, errors
	}
	return obj, nil
}
