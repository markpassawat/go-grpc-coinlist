package errx

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypedError(t *testing.T) {
	var tests = []struct {
		name     string
		err      error
		expected string
	}{
		{
			name: "with only code and message",
			err: &TypedError{
				Code:    "unit_test_code",
				Message: "unit test message",
			},
			expected: `status_code: 0
code: unit_test_code
message: unit test message
`,
		},
		{
			name: "with status code, code, and message",
			err: &TypedError{
				StatusCode: http.StatusUnprocessableEntity,
				Message:    "unit test message",
			},
			expected: `status_code: 422
message: unit test message
`,
		},
		{
			name: "with status code, code, and message, and fields",
			err: &TypedError{
				StatusCode: http.StatusUnprocessableEntity,
				Code:       "unit_test_code",
				Message:    "unit test message",
				Fields: map[string]interface{}{
					"id":     1,
					"field1": "x",
					"field2": true,
				},
			},
			expected: `status_code: 422
code: unit_test_code
message: unit test message
fields:
  field1: x
  field2: true
  id: 1
`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)
			r.Equal(tt.expected, tt.err.Error())
		})
	}
}
