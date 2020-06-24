package validator

type ErrorMessage struct {
	Field   string `json:"field,omitempty" bson:"field,omitempty" gorm:"column:field"`
	Code    string `json:"code,omitempty" bson:"code,omitempty" gorm:"column:code"`
	Message string `json:"message,omitempty" bson:"message,omitempty" gorm:"column:message"`
}
