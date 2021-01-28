package validator

import "github.com/go-playground/validator/v10"

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
	list = append(list, CustomValidate{Fn: CheckIp, Tag: "ip"})
	list = append(list, CustomValidate{Fn: CheckIpV6, Tag: "ipv6"})
	list = append(list, CustomValidate{Fn: CheckDigit, Tag: "digit"})
	list = append(list, CustomValidate{Fn: CheckAbc, Tag: "abc"})
	list = append(list, CustomValidate{Fn: CheckId, Tag: "id"})
	list = append(list, CustomValidate{Fn: CheckCode, Tag: "code"})
	list = append(list, CustomValidate{Fn: CheckCountryCode, Tag: "country_code"})
	list = append(list, CustomValidate{Fn: CheckUsername, Tag: "username"})
	list = append(list, CustomValidate{Fn: CheckPattern, Tag: "regex"})
	return
}

func CheckFax(fl validator.FieldLevel) bool {
	return IsFax(fl.Field().String())
}
func CheckPhone(fl validator.FieldLevel) bool {
	return IsPhone(fl.Field().String())
}
func CheckIp(fl validator.FieldLevel) bool {
	return IsIpAddressV4(fl.Field().String())
}
func CheckIpV6(fl validator.FieldLevel) bool {
	return IsIpAddressV6(fl.Field().String())
}
func CheckDigit(fl validator.FieldLevel) bool {
	return IsDigit(fl.Field().String())
}
func CheckAbc(fl validator.FieldLevel) bool {
	return IsAbc(fl.Field().String())
}
func CheckId(fl validator.FieldLevel) bool {
	return IsCode(fl.Field().String())
}
func CheckCode(fl validator.FieldLevel) bool {
	return IsDashCode(fl.Field().String())
}
func CheckCountryCode(fl validator.FieldLevel) bool {
	return IsCountryCode(fl.Field().String())
}
func CheckUsername(fl validator.FieldLevel) bool {
	return IsUserName(fl.Field().String())
}
func CheckPattern(fl validator.FieldLevel) bool {
	param := fl.Param()
	if pattern, ok := PatternMap[param]; ok {
		return IsValidPattern(pattern, fl.Field().String())
	} else {
		panic("invalid pattern")
	}
}
