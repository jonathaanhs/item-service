package httpkit_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/item-service/pkg/httpkit"
)

func TestNewValidErr(t *testing.T) {
	err := fmt.Errorf("HTTP error 422: some-message")
	errHttpKit := httpkit.NewError(422, "some-message")
	if errHttpKit.Error() != err.Error() {
		t.Fatalf("expected %+v but got %+v", err, errHttpKit)
	}
}

func TestHTTPError(t *testing.T) {
	tests := []struct {
		name               string
		err                error
		expectedStatusCode int
		expectedMessage    string
	}{
		{
			name:               "Should encode error response with provided status code and message",
			err:                httpkit.ErrorResp{StatusCode: http.StatusBadRequest, Message: "Bad Request"},
			expectedStatusCode: http.StatusBadRequest,
			expectedMessage:    "Bad Request",
		},
		{
			name:               "Should encode error response with 500 status code and error message",
			err:                fmt.Errorf("internal server error"),
			expectedStatusCode: http.StatusInternalServerError,
			expectedMessage:    "internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			httpkit.HTTPError(w, tt.err)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatusCode {
				t.Errorf("unexpected status code: got %d, want %d", resp.StatusCode, tt.expectedStatusCode)
			}

			var errResp httpkit.ErrorResp
			err := json.NewDecoder(resp.Body).Decode(&errResp)
			if err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			if errResp.Message != tt.expectedMessage {
				t.Errorf("unexpected error message: got %s, want %s", errResp.Message, tt.expectedMessage)
			}
		})
	}
}
