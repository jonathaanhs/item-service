package httpkit

import (
	"encoding/json"
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
		json.NewEncoder(w).Encode(ErrorResp{StatusCode: httpErr.StatusCode, Message: httpErr.Message})
	} else {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
