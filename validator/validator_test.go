package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Validate(t *testing.T) {
	type testCase struct {
		name      string
		isValid   bool
		hasErrors bool
		Value     string `validate:"required,alphanum"`
		URL       string `validate:"required,url"`
	}

	testCases := []testCase{
		{
			name:    "valid struct",
			Value:   "test",
			URL:     "http://test.com",
			isValid: true,
		},
		{
			name:      "value contains spaces",
			Value:     "test test",
			URL:       "http://test.com",
			hasErrors: true,
		},
		{
			name:      "url not valid",
			Value:     "test",
			URL:       "test.com",
			hasErrors: true,
		},
		{
			name:      "empty value",
			Value:     "",
			URL:       "http://test.com",
			hasErrors: true,
		},
		{
			name:      "empty url",
			Value:     "test",
			URL:       "",
			hasErrors: true,
		},
		{
			name:      "invalid value and url",
			Value:     "test test",
			URL:       "test.com",
			hasErrors: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			v := New()
			require.NotNil(t, v)

			valid := v.Validate(tc)
			assert.Equal(t, tc.isValid, valid)

			if tc.hasErrors {
				assert.NotEmpty(t, v.Errors())
			}
		})
	}
}
