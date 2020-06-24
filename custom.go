package validator

import "gopkg.in/go-playground/validator.v9"

type CustomValidate struct {
	Fn  validator.Func
	Tag string
}

var PatternMap = map[string]string{
	"digit":      "^\\d+$",
	"dash_digit": "^[0-9-]*$",
	"code":       "^\\w*\\d*$",
}

func GetCustomValidateList() (list []CustomValidate) {
	list = append(list, CustomValidate{Fn: CheckFax, Tag: "fax"})
	list = append(list, CustomValidate{Fn: CheckPhone, Tag: "phone"})
	list = append(list, CustomValidate{Fn: CheckDashCode, Tag: "dash_code"})
	list = append(list, CustomValidate{Fn: CheckPattern, Tag: "regex"})
	return
}

func CheckFax(fl validator.FieldLevel) bool {
	return IsFax(fl.Field().String())
}
func CheckPhone(fl validator.FieldLevel) bool {
	return IsPhone(fl.Field().String())
}
func CheckDashCode(fl validator.FieldLevel) bool {
	return IsDashCode(fl.Field().String())
}
func CheckPattern(fl validator.FieldLevel) bool {
	param := fl.Param()
	if pattern, ok := PatternMap[param]; ok {
		return IsValidPattern(pattern, fl.Field().String())
	} else {
		panic("invalid pattern")
	}
}
