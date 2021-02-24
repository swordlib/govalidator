package validators

import (
	"github.com/swordlib/govalidator"
)

// Email return a email ValidatorFunc
// see also https://www.w3.org/TR/2016/REC-html51-20161101/sec-forms.html#email-state-typeemail
func Email() govalidator.ValidatorFunc {
	return func(rule *govalidator.Rule, value interface{}, target interface{}) error {
		vf := Regex("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		err := vf(rule, value, target)
		return err
	}
}
