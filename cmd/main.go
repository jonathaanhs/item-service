package main

import (
	"log"
	"net/http"

	"github.com/tokenomy-assessment/internal"
	"github.com/tokenomy-assessment/internal/controller"
	"github.com/tokenomy-assessment/internal/service"
)

func main() {
	itemSvc := service.NewItemSvc(service.ItemSvcImpl{})

	m := http.NewServeMux()
	controller.NewItemHandler(m, itemSvc)

	if err := internal.StartApp(m, ":8089"); err != nil {
		log.Fatal(err)
	}
}
