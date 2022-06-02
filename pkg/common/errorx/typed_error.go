package errx

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// Internal error is for unhandle error
func Internal() TypedError {
	return TypedError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Something went wrong, please try again",
	}
}

// TypedError contains error details e.g. status code, message
type TypedError struct {
	StatusCode int                    `json:"-" yaml:"status_code"`
	Code       string                 `json:"code,omitempty"  yaml:",omitempty"`
	Message    string                 `json:"message"`
	Fields     map[string]interface{} `json:"fields,omitempty" yaml:",omitempty"`
}

// Error returns the string representation of t.
func (t *TypedError) Error() string {
	fields, err := yaml.Marshal(t)
	if err != nil {
		return fmt.Sprintf("[%d]%s: %s", t.StatusCode, t.Code, t.Message)
	}
	return string(fields)
}
