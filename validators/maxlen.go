package validators

import (
	"errors"
	"fmt"
	"reflect"
	"unicode/utf8"

	"github.com/swordlib/govalidator"
)

// MaxLen return an ValidatorFunc that limit max length of value.
// It panics if value is not Array, Chan, Map, Slice, or String.
func MaxLen(max int) govalidator.ValidatorFunc {
	return func(rule *govalidator.Rule, value interface{}, target interface{}) error {
		v := reflect.Indirect(reflect.ValueOf(value))
		l := 0
		if v.Kind() == reflect.String {
			l = utf8.RuneCountInString(value.(string))
		} else {
			l = v.Len()
		}
		if l > max {
			if rule != nil && rule.Message != "" {
				return errors.New(rule.Message)
			}
			return fmt.Errorf("length(%d) should not greater than %d", l, max)
		}
		return nil
	}
}
