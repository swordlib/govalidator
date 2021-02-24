package validators_test

import (
	"errors"
	"testing"

	"github.com/swordlib/govalidator"
	"github.com/swordlib/govalidator/validators"
)

func TestEmail(t *testing.T) {
	t.Run("Basic usage", func(t *testing.T) {
		type testCase struct {
			Email string
			Want  error
		}
		vf := validators.Email()
		testCases := []testCase{
			{
				Email: "abcd@gmail.com",
				Want:  nil,
			},
			{
				Email: "abcd@gmail.com",
				Want:  nil,
			},
			{
				Email: "job_ms88@163.com",
				Want:  nil,
			},
			{
				Email: "238483@qq.com",
				Want:  nil,
			},
			{
				Email: "238483@mail.qq.com",
				Want:  nil,
			},
			{
				Email: "238483",
				Want:  errors.New("invalid email address"),
			},
		}
		for _, tc := range testCases {
			if got := vf(nil, tc.Email, nil); got != tc.Want {
				if got == nil {
					t.Errorf("want: %q, but got: nil", tc.Want)
				} else if tc.Want == nil {
					t.Errorf("want: nil, but got: %q", got)
				} else if got.Error() != tc.Want.Error() {
					t.Errorf("want: %q, but got: %q", tc.Want, got)
				}
			}
		}
	})
	t.Run("Custom message", func(t *testing.T) {
		vf := validators.Email()
		want := "Invalid email format"
		if got := vf(&govalidator.Rule{
			Message: want,
		}, "abcd", nil); got.Error() != want {
			t.Errorf("want: %q, but got: %q", want, got)
		}
	})
}
