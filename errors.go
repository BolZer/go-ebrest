package ebrest

import (
	"fmt"
	"time"
)

type HttpError struct {
	When         time.Time
	FunctionCall string
	Code         float64
	Message      string
	Arguments    interface{}
}

func (err HttpError) Error() string {
	return fmt.Sprintf("%v. {%v} failed. Error: %v - %v -%v", err.When, err.FunctionCall, err.Code, err.Message, err.Arguments)
}

func newHttpError(function string, code float64, message string, arguments interface{}) error {
	return HttpError{
		When:         time.Now(),
		FunctionCall: function,
		Code:         code,
		Message:      message,
		Arguments:    arguments,
	}
}
