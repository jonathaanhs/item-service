package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/tokenomy-assessment/internal"
	"github.com/tokenomy-assessment/internal/controller"
	"github.com/tokenomy-assessment/internal/service"
)

func main() {
	itemSvc := service.NewItemSvc(service.ItemSvcImpl{})

	m := http.NewServeMux()
	controller.NewItemHandler(m, itemSvc)

	if err := startApp(m, ":8089"); err != nil {
		log.Fatal(err)
	}
}

func startApp(m *http.ServeMux, addr string) error {
	srv := &http.Server{
		Addr:    addr,
		Handler: m,
	}

	go func() {
		log.Println("Server Start ", srv.Addr)
		if err := internal.Start(srv); err != nil {
			log.Fatalf("start: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	if err := internal.Shutdown(srv); err != nil {
		return err
	}

	log.Println("Server exiting")

	return nil
}
