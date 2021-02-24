package validators

import (
	"errors"
	"regexp"

	"github.com/swordlib/govalidator"
)

// Regex return a regular expression ValidatorFunc
func Regex(expr string) govalidator.ValidatorFunc {
	return func(rule *govalidator.Rule, value interface{}, target interface{}) error {
		ok, err := regexp.MatchString(expr, value.(string))
		if err != nil {
			return err
		}
		if !ok {
			if rule != nil && rule.Message != "" {
				return errors.New(rule.Message)
			}
			return errors.New("malformed")
		}
		return nil
	}
}
