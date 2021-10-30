package v9

import v "github.com/core-go/validator"

func NewErrorChecker() *v.ErrorChecker {
	o := NewValidator()
	return v.NewErrorChecker(o.Validate)
}
