package internal

import (
	"context"
	"log"
	"net/http"
	"time"
)

func Shutdown(
	srv *http.Server,
) error {
	log.Printf("Shutdown at %s\n", time.Now().String())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
	return nil
}
