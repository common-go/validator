# Validator
- ErrorMessage
- Validator
- DefaultValidator

## Installation

Please make sure to initialize a Go module before installing core-go/validator:

```shell
go get -u github.com/core-go/validator
```

Import:

```go
import "github.com/core-go/validator"
```

## Details:
#### error_message.go
```go
type ErrorMessage struct {
	Field   string `mapstructure:"field" json:"field,omitempty" gorm:"column:field" bson:"field,omitempty" dynamodbav:"field,omitempty" firestore:"field,omitempty"`
	Code    string `mapstructure:"code" json:"code,omitempty" gorm:"column:code" bson:"code,omitempty" dynamodbav:"code,omitempty" firestore:"code,omitempty"`
	Param   string `mapstructure:"param" json:"param,omitempty" gorm:"column:param" bson:"param,omitempty" dynamodbav:"param,omitempty" firestore:"param,omitempty"`
	Message string `mapstructure:"message" json:"message,omitempty" gorm:"column:message" bson:"message,omitempty" dynamodbav:"message,omitempty" firestore:"message,omitempty"`
}
```

#### validator.go
```go
type Validator interface {
	Validate(ctx context.Context, model interface{}) ([]ErrorMessage, error)
}
```

## Example:
```go
package main

import (
	"context"
	"fmt"
	"github.com/core-go/validator"
	"github.com/core-go/validator/v10"
)

type User struct {
	FirstName string `json:"firstName,omitempty" validate:"required"`
	LastName  string `json:"lastName,omitempty" validate:"required"`
	Email     string `json:"email,omitempty" validate:"omitempty,email"`
}

func main() {
	ctx := context.Background()

	user := User{
		FirstName: "",
		LastName:  "",
		Email:     "peter.parker",
	}

	v := v10.NewValidator()
	errors, _ := v.Validate(ctx, user)
	// Output will be '[{firstName required } {lastName required } {email email }]'
	fmt.Println(errors)
}
```
