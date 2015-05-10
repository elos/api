package hermes

import (
	"fmt"

	"github.com/elos/data"
)

const ErrorKind data.Kind = "error"

type Error struct {
	Status           int    `json:"status"`
	Code             int    `json:"code"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developer_message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("[hermes]: error %d: %s", e.Status, e.Message)
}

var (
	ErrGeneric       = &Error{400, 400, "this is the generic error", "complain to nick about returning the generic error"}
	ErrNotFound      = &Error{400, 400, "data not found", "perhaps you have an incorrect id?"}
	ErrUnauthorized  = &Error{401, 401, "unauthorized", "hit /authenticate first"}
	ErrMalformedData = &Error{400, 400, "malformed data", "json only supported"}
)
