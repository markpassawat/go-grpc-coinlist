package errx

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestErrorx(t *testing.T) {
	typedErr := TypedError{
		StatusCode: http.StatusBadRequest,
		Code:       "invalid_request",
		Message:    "invalid request",
	}
	var tests = []struct {
		name               string
		err                error
		expectedStatusCode int
		expectedFormatted  string
	}{
		{
			name:               "New error",
			err:                New("cannot make request", typedErr),
			expectedStatusCode: http.StatusBadRequest,
			expectedFormatted: `cannot make request
status_code: 400
code: invalid_request
message: invalid request
cannot make request
github.com/Forward-Protocol/APH-event-service/pkg/common/errorx.New
` + `(.*)/errorx.go:\d+
github.com/Forward-Protocol/APH-event-service/pkg/common/errorx.TestErrorx
` + `(.*)
testing.tRunner
` + `(.*)
runtime.goexit
` + `(.*)`,
		},
		{
			name:               "Wrap error",
			err:                Wrap(errors.New("not enough params"), "cannot make request", typedErr),
			expectedStatusCode: http.StatusBadRequest,
			expectedFormatted: `cannot make request: not enough params
status_code: 400
code: invalid_request
message: invalid request
not enough params
github.com/Forward-Protocol/APH-event-service/pkg/common/errorx.TestErrorx
` + `(.*)/errorx_test.go:\d+
testing.tRunner
` + `(.*)
runtime.goexit
` + `(.*)
cannot make request
github.com/Forward-Protocol/APH-event-service/pkg/common/errorx.Wrap
` + `(.*)/errorx.go:\d+
github.com/Forward-Protocol/APH-event-service/pkg/common/errorx.TestErrorx
` + `(.*)/errorx_test.go:\d+
testing.tRunner
` + `(.*)
runtime.goexit
` + `(.*)`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)
			r.Equal(tt.expectedStatusCode, tt.err.(*Errorx).StatusCode())
			r.Regexp(tt.expectedFormatted, fmt.Sprintf("%+v", tt.err))
		})
	}
}
