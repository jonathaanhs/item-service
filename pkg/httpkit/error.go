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
	if httpErr, ok := err.(ErrorResp); ok {
		w.WriteHeader(httpErr.StatusCode)
		json.NewEncoder(w).Encode(ErrorResp{StatusCode: httpErr.StatusCode, Message: httpErr.Message})
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ErrorResp{StatusCode: http.StatusInternalServerError, Message: err.Error()})
}
