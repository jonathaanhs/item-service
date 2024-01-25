package controller

import (
	"encoding/json"
	"net/http"

	"github.com/item-service/internal/service"
	"github.com/item-service/pkg/httpkit"
)

type (
	ItemCntrlImpl struct {
		ItemSvc service.ItemSvc
	}

	Response struct {
		Code int            `json:"code"`
		Data []service.Item `json:"data"`
	}
)

func NewItemHandler(m *http.ServeMux, itemSvc service.ItemSvc) {
	handler := &ItemCntrlImpl{
		ItemSvc: itemSvc,
	}

	m.HandleFunc("/", handler.Get)
}

func (ih *ItemCntrlImpl) Get(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	data, err := ih.ItemSvc.GetItems(idParam)
	if err != nil {
		httpkit.HTTPError(w, err)
		return
	}

	jsonResponse := Response{Code: http.StatusOK, Data: data}
	w.WriteHeader(jsonResponse.Code)
	json.NewEncoder(w).Encode(jsonResponse)
}
