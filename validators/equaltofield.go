package validators

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/swordlib/govalidator"
)

// EqualToField return a ValidatorFunc that validates if fields are equal
func EqualToField(fieldName string) govalidator.ValidatorFunc {
	return func(rule *govalidator.Rule, value interface{}, target interface{}) error {
		t := reflect.Indirect(reflect.ValueOf(target))
		otherField := t.FieldByName(fieldName)
		if !reflect.DeepEqual(value, otherField.Interface()) {
			if rule != nil && rule.Message != "" {
				return errors.New(rule.Message)
			}
			return fmt.Errorf("EqualToField validation error: %q", fieldName)
		}
		return nil
	}
}
