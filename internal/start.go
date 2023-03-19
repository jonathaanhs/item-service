package internal

import (
	"net/http"
)

func Start(
	srv *http.Server,
) (err error) {
	return srv.ListenAndServe()
}
