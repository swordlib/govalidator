package govalidator

import (
	"errors"
	"testing"
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
		HasError bool
		expected error
	}
	cases := []*structTestCase{
		{
			TestName: "Struct Value",
			Data:     testPerson{},
			RulesMap: RulesMap{},
			expected: nil,
		},
		{
			TestName: "Struct Pointer",
			Data:     &testPerson{},
			RulesMap: RulesMap{},
			expected: nil,
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
			expected: errors.New("FirstName is required"),
		},
		{
			TestName: "RequiredStringField",
			Data: &testPerson{
				FirstName: "Alice",
				LastName:  "Smith",
				Country:   "America",
				Age:       18,
			},
			RulesMap: RulesMap{
				"FirstName": {
					{
						Required: true,
					},
				},
			},
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
			expected: errors.New("Age is required"),
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
			expected: nil,
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
			expected: errors.New("Please input your firstname"),
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
						Validator: func(rule *Rule, value interface{}) error {
							if age := value.(int); age < 18 {
								return errors.New("The age of person must be greater than 18 years old")
							}
							return nil
						},
					},
				},
			},
			expected: errors.New("The age of person must be greater than 18 years old"),
		},
	}
	for _, c := range cases {
		t.Run(c.TestName, func(t *testing.T) {
			gv := New(c.RulesMap)
			err := gv.StructFisrt(c.Data)
			if c.expected == nil || err == nil {
				if c.expected == nil && err != nil {
					t.Errorf("Expected nil, but got %q", err)
				} else if c.expected != nil && err == nil {
					t.Errorf("Expected %q, but got nil", c.expected)
				}
			} else if c.expected.Error() != err.Error() {
				t.Errorf("Expected %q, but got %q", c.expected, err)
			}
		})
	}
}
