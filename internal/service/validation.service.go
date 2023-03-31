package service

import "github.com/go-playground/validator/v10"

type ValidationSrv interface {
	Validate(any) error
}

type validationStruct struct{}

func NewValidationService() ValidationSrv {
	return &validationStruct{}
}

func (v *validationStruct) Validate(a any) error {
	return validator.New().Struct(a)
}
