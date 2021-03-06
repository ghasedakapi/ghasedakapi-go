package ghasedakapi

import (
	"fmt"
)

type APIError struct {
	Code  int
	Message string
	Err     error
}
type HTTPError struct {
	Code  int
	Message string
	Err     error
}

func (e *APIError) Error() string {
	return fmt.Sprintf("APIError[%d] : %s", e.Code, e.Message)
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError[%d] : %s", e.Code, e.Message)
}
