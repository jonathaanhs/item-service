package httpkit

import (
	"fmt"
	"net/http"
)

type ErrorResp struct {
	StatusCode int
	Message    string
}

func (e ErrorResp) Error() string {
	return fmt.Sprintf("HTTP error %d: %s", e.StatusCode, e.Message)
}

func NewError(code int, message string) error {
	return ErrorResp{StatusCode: code, Message: message}
}

func HTTPError(w http.ResponseWriter, err error) {
	httpErr, ok := err.(ErrorResp)
	if ok {
		http.Error(w, httpErr.Message, httpErr.StatusCode)
	} else {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
