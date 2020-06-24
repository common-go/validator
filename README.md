# Validator
- ErrorMessage
- Validator
- DefaultValidator

## Installation

Please make sure to initialize a Go module before installing common-go/validator:

```shell
go get -u github.com/common-go/validator
```

Import:

```go
import "github.com/common-go/validator"
```

## Details:
#### error_message.go
```go
type ErrorMessage struct {
	Field   string `json:"field,omitempty" bson:"field,omitempty" gorm:"column:field"`
	Code    string `json:"code,omitempty" bson:"code,omitempty" gorm:"column:code"`
	Message string `json:"message,omitempty" bson:"message,omitempty" gorm:"column:message"`
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
	"github.com/common-go/validator"
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

	v := validator.NewDefaultValidator()
	errors, _ := v.Validate(ctx, user)
	// Output will be '[{firstName required } {lastName required } {email email }]'
	fmt.Println(errors)
}
```
