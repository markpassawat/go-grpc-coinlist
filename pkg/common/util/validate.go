package util

import (
	"fmt"

	errx "github.com/Forward-Protocol/APH-event-service/pkg/common/errorx"
	"gopkg.in/go-playground/validator.v9"
)

func Valid(s interface{}) error {
	if err := validator.New().Struct(s); err != nil {
		return errx.Wrap(err, "validate struct", errx.Internal())
	}
	return nil
}

func MustValid(s interface{}) {
	if err := Valid(s); err != nil {
		fmt.Println(err)
	}
}
