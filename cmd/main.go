package main

import (
	"log"
	"net/http"

	"github.com/item-service/internal"
	"github.com/item-service/internal/controller"
	"github.com/item-service/internal/service"
)

func main() {
	itemSvc := service.NewItemSvc(service.ItemSvcImpl{})

	m := http.NewServeMux()
	controller.NewItemHandler(m, itemSvc)

	if err := internal.StartApp(m, ":8080"); err != nil {
		log.Fatal(err)
	}
}
