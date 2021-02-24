package validators_test

import (
	"errors"
	"testing"

	"github.com/swordlib/govalidator"
	"github.com/swordlib/govalidator/validators"
)

func TestMaxLen(t *testing.T) {
	t.Run("Basic usage", func(t *testing.T) {
		type testCase struct {
			Value     interface{}
			MaxLength int
			Want      error
		}
		testCases := []testCase{
			{
				Value:     "abcd",
				MaxLength: 10,
				Want:      nil,
			},
			{
				Value:     "abcd",
				MaxLength: 2,
				Want:      errors.New("length should not greater than 2"),
			},
			{
				Value:     []int{1, 2, 3, 4},
				MaxLength: 5,
				Want:      nil,
			},
			{
				Value:     [4]int{1, 2, 3, 4},
				MaxLength: 3,
				Want:      errors.New("length should not greater than 3"),
			},
		}
		for _, tc := range testCases {
			if got := validators.MaxLen(tc.MaxLength)(nil, tc.Value, nil); got != tc.Want {
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
		vf := validators.MaxLen(2)
		want := "too long"
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
		vf := validators.MaxLen(8)
		vf(nil, 8, nil)
	})
}
