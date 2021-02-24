package validators

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/swordlib/govalidator"
)

// MaxLen return an ValidatorFunc that limit max length of value
// It panics if value is not Array, Chan, Map, Slice, or String
func MaxLen(max int) govalidator.ValidatorFunc {
	return func(rule *govalidator.Rule, value interface{}, target interface{}) error {
		v := reflect.Indirect(reflect.ValueOf(value))
		if v.Len() > max {
			if rule != nil && rule.Message != "" {
				return errors.New(rule.Message)
			}
			return fmt.Errorf("length should not greater than %d", max)
		}
		return nil
	}
}
