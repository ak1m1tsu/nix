package validator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validate(t *testing.T) {
	type testCase struct {
		name        string
		expectedErr error
		Value       string `validate:"required,alphanum"`
		URL         string `validate:"required,url"`
	}

	testCases := []testCase{
		{
			name:  "valid struct",
			Value: "test",
			URL:   "http://test.com",
		},
		{
			name:        "value contains spaces",
			Value:       "test test",
			URL:         "http://test.com",
			expectedErr: errors.New("Value can only contain alphanumeric characters"),
		},
		{
			name:        "url not valid",
			Value:       "test",
			URL:         "test.com",
			expectedErr: errors.New("URL must be a valid URL"),
		},
		{
			name:        "empty value",
			Value:       "",
			URL:         "http://test.com",
			expectedErr: errors.New("Value is a required field"),
		},
		{
			name:        "empty url",
			Value:       "test",
			URL:         "",
			expectedErr: errors.New("URL is a required field"),
		},
		{
			name:        "invalid value and url",
			Value:       "test test",
			URL:         "test.com",
			expectedErr: errors.New("Value can only contain alphanumeric characters, URL must be a valid URL"),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := Validate(tc)

			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
