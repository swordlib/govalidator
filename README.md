# govalidator
Go Validator

It's unstable now. I will realse v1.0.0, when it finish

## document

[Access api and example document](https://pkg.go.dev/github.com/swordlib/govalidator)

## basic usage

```go
package govalidator_test

import (
	"errors"
	"fmt"

	"github.com/swordlib/govalidator"
)

type person struct {
	FirstName string
}

func Example() {
	gv := govalidator.New(govalidator.RulesMap{
		"FirstName": {
			{
				Required: true,
			},
			{
				Validator: func(rule *govalidator.Rule, value interface{}) error {

					if v := value.(string); v != "Alice" {
						return errors.New("FirstName must be Alice")
					}
					return nil
				},
			},
		},
	})
	fmt.Println(gv.StructFirst(&person{
		FirstName: "",
	}))
	fmt.Println(gv.StructFirst(&person{
		FirstName: "Bob",
	}))
	fmt.Println(gv.StructFirst(&person{
		FirstName: "Alice",
	}))

	custom := govalidator.New(govalidator.RulesMap{
		"FirstName": {
			{
				Required: true,
				Message:  "Please input your firstName",
			},
		},
	})
	fmt.Println(custom.StructFirst(&person{}))
	// Output:
	// FirstName is required
	// FirstName must be Alice
	// <nil>
	// Please input your firstName
}
```