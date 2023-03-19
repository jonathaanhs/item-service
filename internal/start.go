package internal

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func StartApp(m *http.ServeMux, addr string) error {
	srv := &http.Server{
		Addr:    addr,
		Handler: m,
	}

	go func() {
		log.Println("Server Start ", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("start: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	if err := Shutdown(srv); err != nil {
		return err
	}

	log.Println("Server exiting")

	return nil
}
