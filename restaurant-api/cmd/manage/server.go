package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) serve() error {

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", app.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return srv.ListenAndServe()

}
