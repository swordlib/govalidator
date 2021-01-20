package govalidator

import (
	"errors"
	"fmt"
	"reflect"
)

type Rule struct {
	Required  bool
	Validator ValidateFunc
	Message   string
}

func (r *Rule) validate(value reflect.Value, name string) (err error) {
	// custom error message
	defer (func() {
		if err != nil && r.Validator == nil && r.Message != "" {
			err = errors.New(r.Message)
		}
	})()

	if r.Required && value.IsZero() {
		err = fmt.Errorf("%s is required", name)
		return
	}

	if r.Validator != nil {
		if err = r.Validator(r, value.Interface()); err != nil {
			return
		}
	}
	return
}

type Rules []*Rule

func (rs Rules) validate(value reflect.Value, name string) error {
	for _, rule := range rs {
		if err := rule.validate(value, name); err != nil {
			return err
		}
	}
	return nil
}

type RulesMap map[string]Rules

// ValidatorOptions It's not used until now, just reserves for future
type ValidatorOptions struct {
}

// Validator is a validation program for go
type Validator struct {
	options *ValidatorOptions
	rules   RulesMap
}

// StructFirst Validate a struct and stop when it encounter the first error
func (v *Validator) StructFirst(value interface{}) error {
	rv := reflect.Indirect(reflect.ValueOf(value))
	for fieldName, fieldRules := range v.rules {
		fv := rv.FieldByName(fieldName)
		if err := fieldRules.validate(fv, fieldName); err != nil {
			return err
		}
	}
	return nil
}

// ValidateFunc is a custom validator type, alias of func(rule *Rule, value interface{}) error
type ValidateFunc func(rule *Rule, value interface{}) error

// New create a new validator
func New(rules RulesMap, varoptions ...*ValidatorOptions) *Validator {
	var v *Validator
	if len(varoptions) > 0 && varoptions[0] != nil {
		v = &Validator{
			options: varoptions[0],
		}
	}
	v = &Validator{
		options: &ValidatorOptions{},
	}
	v.rules = rules
	return v
}
