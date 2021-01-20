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
