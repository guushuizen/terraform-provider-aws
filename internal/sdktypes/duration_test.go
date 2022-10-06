package sdktypes

import (
	"regexp"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	cases := []struct {
		val           string
		expectNull    bool
		expectedValue time.Duration
		expectErr     bool
	}{
		{
			val:           "1h2m3s",
			expectNull:    false,
			expectedValue: 1*time.Hour + 2*time.Minute + 3*time.Second,
		},
		{
			val:           "",
			expectNull:    true,
			expectedValue: 0,
		},
		{
			val:           "A",
			expectNull:    false,
			expectedValue: 0,
			expectErr:     true,
		},
	}

	for i, tc := range cases {
		v := Duration(tc.val)

		if null := v.IsNull(); null != tc.expectNull {
			t.Fatalf("expected test case %d IsNull to return %t, got %t", i, null, tc.expectNull)
		}

		value, null, err := v.Value()
		if value != tc.expectedValue {
			t.Fatalf("expected test case %d Value to be %s, got %s", i, tc.expectedValue, value)
		}
		if null != tc.expectNull {
			t.Fatalf("expected test case %d Value null flag to be %t, got %t", i, tc.expectNull, null)
		}
		if tc.expectErr == false && err != nil {
			t.Fatalf("expected test case %d to succeed, got error %s", i, err)
		}
		if tc.expectErr && err == nil {
			t.Fatalf("expected test case %d to have error but had none", i)
		}
	}
}

func TestValidationDuration(t *testing.T) {
	runTestCases(t, map[string]testCase{
		"valid": {
			val: "1h2m3s",
			f:   ValidateDuration,
		},
		"invalid": {
			val:             "A",
			f:               ValidateDuration,
			expectedSummary: regexp.MustCompile(`^Invalid value$`),
			expectedDetail:  regexp.MustCompile(`time: invalid duration "A"`),
		},
		"wrong type": {
			val:             1,
			f:               ValidateDuration,
			expectedSummary: regexp.MustCompile(`^Invalid value type$`),
			expectedDetail:  regexp.MustCompile(`Expected type to be string`),
		},
	})
}
