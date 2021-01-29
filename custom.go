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
	list = append(list, CustomValidate{Fn: CheckEmail, Tag: "email"})
	list = append(list, CustomValidate{Fn: CheckUrl, Tag: "url"})
	list = append(list, CustomValidate{Fn: CheckUri, Tag: "uri"})
	list = append(list, CustomValidate{Fn: CheckFax, Tag: "fax"})
	list = append(list, CustomValidate{Fn: CheckPhone, Tag: "phone"})
	list = append(list, CustomValidate{Fn: CheckIp, Tag: "ip"})
	list = append(list, CustomValidate{Fn: CheckIpV4, Tag: "ipv4"})
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
func CheckEmail(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return true
	}
	return IsEmail(s)
}
func CheckUrl(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return true
	}
	return IsUrl(s)
}
func CheckUri(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return true
	}
	return IsUri(s)
}
func CheckFax(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return true
	}
	return IsFax(s)
}
func CheckPhone(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return true
	}
	return IsPhone(s)
}
func CheckIp(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return true
	}
	return IsIpAddress(s)
}
func CheckIpV4(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return true
	}
	return IsIpAddressV4(s)
}
func CheckIpV6(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return true
	}
	return IsIpAddressV6(s)
}
func CheckDigit(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return true
	}
	return IsDigit(s)
}
func CheckAbc(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return true
	}
	return IsAbc(s)
}
func CheckId(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return true
	}
	return IsCode(s)
}
func CheckCode(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return true
	}
	return IsDashCode(s)
}
func CheckCountryCode(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return true
	}
	return IsCountryCode(s)
}
func CheckUsername(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return true
	}
	return IsUserName(s)
}
func CheckPattern(fl validator.FieldLevel) bool {
	param := fl.Param()
	if pattern, ok := PatternMap[param]; ok {
		return IsValidPattern(pattern, fl.Field().String())
	} else {
		panic("invalid pattern")
	}
}
