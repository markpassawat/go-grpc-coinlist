package errx

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// Errorx is error with extened features
type Errorx struct {
	terr TypedError
	err  error
}

// Unwrap unwraps wrapped error, if any.
func (e *Errorx) Unwrap() error {
	return errors.Cause(e.err)
}

// Typed returns typed error.
func (e *Errorx) Typed() TypedError {
	return e.terr
}

// StatusCode returns HTTP status code-like status code.
func (e *Errorx) StatusCode() int {
	return e.terr.StatusCode
}

// Error returns string representation of e. This should not be used by
// machine but rather can be read and understand quickly by human.
func (e *Errorx) Error() string {
	if e.err == nil {
		return ""
	}
	return e.err.Error()
}

// Format writes string representation of e based on the given verb format to s.
//
// When formatting with +v verb it will include code, message, fields, and stack trace.
//
// For other verbs it will write e.Error() to s.
func (e *Errorx) Format(s fmt.State, verb rune) {
	if verb == 'v' && s.Flag('+') {
		if e.err != nil {
			io.WriteString(s, e.err.Error())
			io.WriteString(s, "\n")
		}

		io.WriteString(s, e.terr.Error())

		if e.err != nil {
			errWithFormat, ok := e.err.(interface {
				Format(fmt.State, rune)
			})
			if ok {
				errWithFormat.Format(s, verb)
			} else {
				io.WriteString(s, e.err.Error())
			}
		} else {
			io.WriteString(s, e.Error())
		}
	} else {
		io.WriteString(s, e.Error())
	}
}

// New creates error with optional additional context and a required typed error.
// Typed error can be created only in this package, making it explicitly defined before
// it can be used. Thus increases overall code and API quality.
func New(ctx string, terr TypedError) error {
	return &Errorx{
		terr: terr,
		err:  errors.New(ctx),
	}
}

func Newf(ctxFormat string, terr TypedError, args ...interface{}) error {
	return &Errorx{
		terr: terr,
		err:  errors.New(fmt.Sprintf(ctxFormat, args...)),
	}
}

// Wrap wrapps err with optional additional context and a required typed error.
// Typed error can be created only in this package, making it explicitly defined before
// it can be used. Thus increases overall code and API quality.
func Wrap(err error, ctx string, terr TypedError) error {
	return &Errorx{
		terr: terr,
		err:  errors.Wrap(err, ctx),
	}
}

// Wrapf is like Wrap but with format string ctxFormat and its arguments args.
func Wrapf(err error, ctxFormat string, terr TypedError, args ...interface{}) error {
	return &Errorx{
		terr: terr,
		err:  errors.Wrapf(err, ctxFormat, args...),
	}
}
