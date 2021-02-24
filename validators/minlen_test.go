package validators_test

import (
	"errors"
	"testing"

	"github.com/swordlib/govalidator"
	"github.com/swordlib/govalidator/validators"
)

func TestMinLen(t *testing.T) {
	t.Run("Basic usage", func(t *testing.T) {
		type testCase struct {
			Value     interface{}
			MinLength int
			Want      error
		}
		testCases := []testCase{
			{
				Value:     "abcd",
				MinLength: 2,
				Want:      nil,
			},
			{
				Value:     "abcd",
				MinLength: 6,
				Want:      errors.New("length(4) must be greater than 6"),
			},
			{
				Value:     "支持中文",
				MinLength: 6,
				Want:      errors.New("length(4) must be greater than 6"),
			},
			{
				Value:     []int{1, 2, 3, 4},
				MinLength: 4,
				Want:      nil,
			},
			{
				Value:     [4]int{1, 2, 3, 4},
				MinLength: 6,
				Want:      errors.New("length(4) must be greater than 6"),
			},
		}
		for _, tc := range testCases {
			if got := validators.MinLen(tc.MinLength)(nil, tc.Value, nil); got != tc.Want {
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
		vf := validators.MinLen(6)
		want := "too short"
		if got := vf(&govalidator.Rule{
			Message: want,
		}, "abcd", nil); got.Error() != want {
			t.Errorf("want: %q, but got: %q", want, got)
		}
	})
	t.Run("Using with value of unsupported type", func(t *testing.T) {
		defer (func() {
			if panicMsg := recover(); panicMsg == nil {
				t.Errorf("want a panic, but got nil")
			}
		})()
		vf := validators.MinLen(8)
		vf(nil, 8, nil)
	})
}
