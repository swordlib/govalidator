package govalidator

import (
	"fmt"
	"reflect"
)

type Rule struct {
	Required  bool
	Validator *ValidateFunc
	Message   string
}

func (r *Rule) validate(value reflect.Value, name string) error {
	if r.Required && value.IsZero() {
		return fmt.Errorf("%s is required", name)
	}
	return nil
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
type Validator struct {
	options *ValidatorOptions
	rules   RulesMap
}

func (v *Validator) StructFisrt(value interface{}) error {
	rv := reflect.Indirect(reflect.ValueOf(value))
	for fieldName, fieldRules := range v.rules {
		fv := rv.FieldByName(fieldName)
		if err := fieldRules.validate(fv, fieldName); err != nil {
			return err
		}
	}
	return nil
}

type ValidateFunc func(rule, value interface{}) error

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
