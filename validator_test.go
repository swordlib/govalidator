package govalidator_test

import (
	"errors"
	"testing"

	. "github.com/swordlib/govalidator"
)

type testPerson struct {
	FirstName string
	LastName  string
	Country   string
	Age       int
}

func TestNew(t *testing.T) {
	t.Run("NewWithoutOptions", func(t *testing.T) {
		New(RulesMap{})
	})
	t.Run("NewWithOptions", func(t *testing.T) {
		New(RulesMap{}, &ValidatorOptions{})
	})
}

func TestStructFirst(t *testing.T) {
	type structTestCase struct {
		TestName string
		Data     interface{}
		RulesMap RulesMap
		want     error
	}
	cases := []*structTestCase{
		{
			TestName: "Struct Value",
			Data:     testPerson{},
			RulesMap: RulesMap{},
			want:     nil,
		},
		{
			TestName: "Struct Pointer",
			Data:     &testPerson{},
			RulesMap: RulesMap{},
			want:     nil,
		},
		{
			TestName: "EmptyRequiredStringField",
			Data: &testPerson{
				FirstName: "",
			},
			RulesMap: RulesMap{
				"FirstName": {
					{
						Required: true,
					},
				},
			},
			want: errors.New("FirstName is required"),
		},
		{
			TestName: "RequiredStringField",
			Data: &testPerson{
				FirstName: "Alice",
			},
			RulesMap: RulesMap{
				"FirstName": {
					{
						Required: true,
					},
				},
			},
			want: nil,
		},
		{
			TestName: "EmptyRequiredIntField",
			Data: &testPerson{
				FirstName: "Alice",
			},
			RulesMap: RulesMap{
				"Age": {
					{
						Required: true,
					},
				},
			},
			want: errors.New("Age is required"),
		},
		{
			TestName: "RequiredIntField",
			Data: &testPerson{
				FirstName: "Alice",
				Age:       18,
			},
			RulesMap: RulesMap{
				"Age": {
					{
						Required: true,
					},
				},
			},
			want: nil,
		},
		{
			TestName: "CustomErrorMessage",
			Data: &testPerson{
				FirstName: "",
			},
			RulesMap: RulesMap{
				"FirstName": {
					{
						Required: true,
						Message:  "Please input your firstname",
					},
				},
			},
			want: errors.New("Please input your firstname"),
		},
		{
			TestName: "CustomValidator",
			Data: &testPerson{
				FirstName: "Alice",
				Age:       16,
			},
			RulesMap: RulesMap{
				"FirstName": {
					{
						Required: true,
					},
				},
				"Age": {
					{
						Validator: func(rule *Rule, value interface{}, target interface{}) error {
							if age := value.(int); age < 18 {
								return errors.New("The age of person must be greater than 18 years old")
							}
							return nil
						},
					},
				},
			},
			want: errors.New("The age of person must be greater than 18 years old"),
		},
		{
			TestName: "MessageShouldNotOverrideCustomValidatorError",
			Data: &testPerson{
				FirstName: "Alice",
			},
			RulesMap: RulesMap{
				"FirstName": {
					{
						Required: true,
						Message:  "Please input your firstname",
						Validator: func(rule *Rule, value interface{}, target interface{}) error {
							return errors.New("The name Alice has been occupied")
						},
					},
				},
			},
			want: errors.New("The name Alice has been occupied"),
		},
	}
	for _, c := range cases {
		t.Run(c.TestName, func(t *testing.T) {
			gv := New(c.RulesMap)
			err := gv.StructFirst(c.Data)
			if c.want == nil || err == nil {
				if c.want == nil && err != nil {
					t.Errorf("want: nil, but got: %q", err)
				} else if c.want != nil && err == nil {
					t.Errorf("want: %q, but got: nil", c.want)
				}
			} else if c.want.Error() != err.Error() {
				t.Errorf("want: %q, but got: %q", c.want, err)
			}
		})
	}
}

func TestStructWithOtherThing(t *testing.T) {
	defer (func() {
		if err := recover(); err != nil {
			want := "value must be a struct"
			if err.(string) != want {
				t.Errorf("want: %q, but got %q", want, err)
			}
		} else {
			t.Error("Call StructFirst with other than struct should panic")
		}
	})()
	gv := New(RulesMap{})
	gv.StructFirst(8)
	gv.StructFirst("testing")
	gv.StructFirst(9.0)
	gv.StructFirst([]int{8, 8, 8})
}

func TestValidatingNotPresentStructField(t *testing.T) {
	defer (func() {
		if err := recover(); err != nil {
			want := `Struct field: "NotPresent" is not present`
			if err.(string) != want {
				t.Errorf("want: %q, but got %q", want, err)
			}
		} else {
			t.Error("Validating not present struct field should panic")
		}
	})()
	New(RulesMap{
		"NotPresent": {
			{
				Required: true,
			},
		},
	}).StructFirst(&testPerson{})
}
