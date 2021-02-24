package validators

import (
	"github.com/swordlib/govalidator"
)

// Email return a email ValidatorFunc
func Email() govalidator.ValidatorFunc {
	return func(rule *govalidator.Rule, value interface{}, target interface{}) error {
		return nil
	}
}
