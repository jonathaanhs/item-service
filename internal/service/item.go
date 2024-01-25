package service

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/item-service/pkg/httpkit"
)

type (
	ItemSvc interface {
		GetItems(ids string) ([]Item, error)
	}

	ItemSvcImpl struct{}

	Item struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
)

var (
	ItemData = map[int]Item{
		1: {1, "A"},
		2: {2, "B"},
		3: {3, "C"},
	}
)

func NewItemSvc(impl ItemSvcImpl) ItemSvc {
	return &impl
}

func (t *ItemSvcImpl) GetItems(ids string) ([]Item, error) {
	if ids != "" {
		return t.getItemsByID(ids)
	}

	return t.getAllItems(), nil
}

func (t *ItemSvcImpl) getAllItems() []Item {
	var items []Item
	for _, v := range ItemData {
		items = append(items, Item{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].ID < items[j].ID
	})
	return items
}

func (t *ItemSvcImpl) getItemsByID(ids string) ([]Item, error) {
	var items []Item
	for _, id := range strings.Split(ids, ",") {
		if id == "" {
			continue
		}
		itemID, err := strconv.Atoi(id)
		if err != nil {
			return nil, httpkit.NewError(http.StatusBadRequest, fmt.Sprintf("invalid or empty ID: %q", id))
		}
		if _, ok := ItemData[itemID]; ok {
			items = append(items, ItemData[itemID])
		}
	}

	if len(items) == 0 {
		return nil, httpkit.NewError(http.StatusNotFound, fmt.Sprintf("resource with ID %s not exist", ids))
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].ID < items[j].ID
	})

	return items, nil
}
