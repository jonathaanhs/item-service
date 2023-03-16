package service_test

import (
	"net/http"
	"testing"

	"github.com/tokenomy-assessment/internal/service"
	"github.com/tokenomy-assessment/pkg/httpkit"
)

func TestItemSvc_GetItems(t *testing.T) {
	type itemSvcTest struct {
		name      string
		ids       string
		expected  []service.Item
		expectErr error
	}

	tests := []itemSvcTest{
		{
			name: "get all items",
			ids:  "",
			expected: []service.Item{
				{1, "A"},
				{2, "B"},
				{3, "C"},
			},
			expectErr: nil,
		},
		{
			name: "get one item by id",
			ids:  "1",
			expected: []service.Item{
				{1, "A"},
			},
			expectErr: nil,
		},
		{
			name: "get multiple items by id",
			ids:  "1,3",
			expected: []service.Item{
				{1, "A"},
				{3, "C"},
			},
			expectErr: nil,
		},
		{
			name:      "invalid or empty ID",
			ids:       "1,4,10,x,y,z",
			expected:  []service.Item{},
			expectErr: httpkit.NewError(http.StatusBadRequest, "invalid or empty ID: \"x\""),
		},
		{
			name:      "resource with ID not exist",
			ids:       "4,5,6",
			expected:  []service.Item{},
			expectErr: httpkit.NewError(http.StatusNotFound, "resource with ID 4,5,6 not exist"),
		},
		{
			name: "get multiple items with empty first ID",
			ids:  ",2,3",
			expected: []service.Item{
				{2, "B"},
				{3, "C"},
			},
			expectErr: nil,
		},
		{
			name:      "invalid single id",
			ids:       "xxx",
			expected:  []service.Item{},
			expectErr: httpkit.NewError(http.StatusBadRequest, "invalid or empty ID: \"xxx\""),
		},
		{
			name:      "resource with single ID not exist",
			ids:       "4",
			expected:  []service.Item{},
			expectErr: httpkit.NewError(http.StatusNotFound, "resource with ID 4 not exist"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			itemSvcImpl := service.ItemSvcImpl{}
			itemSvc := service.NewItemSvc(itemSvcImpl)

			items, err := itemSvc.GetItems(tc.ids)
			if tc.expectErr != err {
				t.Errorf("expected error %v, but got %v", tc.expectErr, err)
			}

			if len(items) != len(tc.expected) {
				t.Errorf("expected length of items to be %d, but got %d", len(tc.expected), len(items))
			}

			for i, item := range items {
				if item.ID != tc.expected[i].ID {
					t.Errorf("expected item ID to be %d, but got %d", tc.expected[i].ID, item.ID)
				}

				if item.Name != tc.expected[i].Name {
					t.Errorf("expected item name to be %s, but got %s", tc.expected[i].Name, item.Name)
				}
			}
		})
	}
}
