package controller

import (
	"github.com/asaskevich/govalidator"
)

type CustomValidator struct{}

func (cv *CustomValidator) Validate(i any) error {
	_, err := govalidator.ValidateStruct(i)
	if err != nil {
		return err
	}
	return nil
}
