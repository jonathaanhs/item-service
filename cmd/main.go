package main

import (
	"log"
	"net/http"

	"github.com/tokenomy-assessment/internal/controller"
	"github.com/tokenomy-assessment/internal/service"
)

func main() {
	itemSvc := service.NewItemSvc(service.ItemSvcImpl{})

	m := http.NewServeMux()
	controller.NewItemHandler(m, itemSvc)

	srv := &http.Server{
		Addr:    ":8089",
		Handler: m,
	}

	log.Println("Server Start ", ":8089")

	srv.ListenAndServe()
}
