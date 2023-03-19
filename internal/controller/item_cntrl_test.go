package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tokenomy-assessment/internal/controller"
	"github.com/tokenomy-assessment/internal/service"
	"github.com/tokenomy-assessment/pkg/httpkit"
)

type mockItemSvc struct{}

func (m *mockItemSvc) GetItems(ids string) ([]service.Item, error) {
	if ids == "x" {
		return []service.Item{}, httpkit.NewError(http.StatusBadRequest, "invalid or empty ID: \"x\"")
	}

	return []service.Item{
		{ID: 1, Name: "A"},
		{ID: 2, Name: "B"},
		{ID: 3, Name: "C"},
	}, nil
}

func TestItemCntrlImpl_Get(t *testing.T) {
	tests := []struct {
		name             string
		url              string
		expectedResponse interface{}
	}{
		{
			name: "should return list item",
			url:  "/?id=",
			expectedResponse: controller.Response{
				Code: http.StatusOK,
				Data: []service.Item{
					{ID: 1, Name: "A"},
					{ID: 2, Name: "B"},
					{ID: 3, Name: "C"},
				},
			},
		},
		{
			name:             "should return error",
			url:              "/?id=x",
			expectedResponse: httpkit.NewError(http.StatusBadRequest, "invalid or empty ID: \"x\""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := &mockItemSvc{}
			router := http.NewServeMux()
			controller.NewItemHandler(router, mockSvc)
			request, err := http.NewRequest(http.MethodGet, tt.url, http.NoBody)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()

			ic := &controller.ItemCntrlImpl{
				ItemSvc: mockSvc,
			}
			router.HandleFunc(tt.url, ic.Get)
			router.ServeHTTP(rr, request)

			resultByte, err := json.Marshal(tt.expectedResponse)
			if err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(string(resultByte)) {
				t.Errorf("unexpected response body:\n\ngot:\n%s\n\nwant:\n%s", strings.TrimSpace(rr.Body.String()), strings.TrimSpace(string(resultByte)))
			}
		})
	}
}
