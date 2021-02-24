package validators_test

import (
	"testing"

	"github.com/swordlib/govalidator"
	"github.com/swordlib/govalidator/validators"
)

func TestRegex(t *testing.T) {
	t.Run("Basic Usage", func(t *testing.T) {
		vf := validators.Regex(`foo.`)
		if got := vf(nil, "food", nil); got != nil {
			t.Errorf("want: nil, but got: %q", got)
		}
		if got := vf(nil, "fog", nil); got == nil || got.Error() != "malformed" {
			t.Errorf("want: %q, but got: %q", "malformed", got)
		}
	})
	t.Run("Custom Message", func(t *testing.T) {
		vf := validators.Regex(`\d+`)
		want := "invalid format"
		if got := vf(&govalidator.Rule{
			Message: want,
		}, "text", nil); got.Error() != want {
			t.Errorf("want: %q, but got: %q", want, got)
		}
	})
	t.Run("Invalid regular expression", func(t *testing.T) {
		vf := validators.Regex(`(`)
		if got := vf(nil, "text", nil); got == nil {
			t.Error("want an error, but got nil")
		}
	})
}
