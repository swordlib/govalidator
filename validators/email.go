package validators

import (
	"errors"

	"github.com/swordlib/govalidator"
)

// Email return an email ValidatorFunc
// see also https://www.w3.org/TR/2016/REC-html51-20161101/sec-forms.html#email-state-typeemail
func Email() govalidator.ValidatorFunc {
	return func(rule *govalidator.Rule, value interface{}, target interface{}) error {
		vf := Regex("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		err := vf(rule, value, target)
		if err != nil && err.Error() == "malformed" {
			return errors.New("invalid email address")
		}
		return err
	}
}
