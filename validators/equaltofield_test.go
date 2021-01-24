package validators_test

import (
	"testing"

	"github.com/swordlib/govalidator/validators"
)

type testAccount struct {
	Password     string
	PasswordCopy string
}

func TestEqualToField(t *testing.T) {
	errorTa := &testAccount{
		Password:     "a",
		PasswordCopy: "b",
	}
	validator := validators.EqualToField("Password")
	want := `EqualToField validation error: "Password"`
	got := validator(nil, errorTa.PasswordCopy, errorTa)
	if got == nil {
		t.Fatalf("Two fields that are't equal should got an error, but got nil")
	}
	if got.Error() != want {
		t.Errorf("want: %q, but got: %q", want, got)
	}

	correctTa := &testAccount{
		Password:     "a",
		PasswordCopy: "a",
	}
	got = validator(nil, correctTa.PasswordCopy, correctTa)
	if got != nil {
		t.Errorf("want: nil, but got: %q", got)
	}
}

func TestCustomMessage(t *testing.T) {
	ta := &testAccount{
		Password:     "a",
		PasswordCopy: "b",
	}
	want := "Two password must be same"
	validator := validators.EqualToField("Password", want)
	got := validator(nil, ta.PasswordCopy, ta)
	if got == nil {
		t.Fatalf("want: %q, but got: nil", want)
	}
	if got.Error() != want {
		t.Errorf("want: %q, but got: %q", want, got)
	}
}

func TestNotPresentStructField(t *testing.T) {
	defer (func() {
		if panicMsg := recover(); panicMsg == nil {
			t.Errorf("Expect a panic when use with not present struct field.")
		}
	})()

	ta := &testAccount{
		Password:     "a",
		PasswordCopy: "b",
	}
	validator := validators.EqualToField("NotPresentPassword")
	validator(nil, ta.PasswordCopy, ta)
}
