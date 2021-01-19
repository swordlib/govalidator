package govalidator

import (
	"errors"
	"testing"
)

type testPerson struct {
	FirstName string
	LastName  string
	Country   string
}

func TestStructFirst(t *testing.T) {
	type structTestCase struct {
		TestName string
		Data     interface{}
		RulesMap RulesMap
		HasError bool
		expected error
	}
	cases := []*structTestCase{
		{
			TestName: "EmptyRequiredFieldShouldError",
			Data: &testPerson{
				"", "Smith", "America",
			},
			RulesMap: RulesMap{
				"FirstName": {
					{
						Required: true,
					},
				},
			},
			expected: errors.New("FirstName is required"),
		},
		{
			TestName: "NotEmptyRequiredFieldShouldPass",
			Data: &testPerson{
				"alice", "Smith", "America",
			},
			RulesMap: RulesMap{
				"FirstName": {
					{
						Required: true,
					},
				},
			},
			expected: nil,
		},
	}
	for _, c := range cases {
		t.Run(c.TestName, func(t *testing.T) {
			gv := New(c.RulesMap)
			err := gv.StructFisrt(c.Data)
			if c.expected == nil || err == nil {
				if c.expected == nil && err != nil {
					t.Errorf("expected nil, but got %q", err)
				} else if c.expected != nil && err == nil {
					t.Errorf("expected %q, but got nil", c.expected)
				}
			} else if c.expected.Error() != err.Error() {
				t.Errorf("expected %q, but got %q", c.expected, err)
			}
		})
	}
}
