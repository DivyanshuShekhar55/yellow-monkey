package main

import (
	"net/http"

	"github.com/DivyanshuShekhar55/yellow-monkey/backend/internal/routes"
)

type application struct {
	conf config
}

type config struct {
	addr string
	//db   dbConfig
}

type dbConfig struct {
	addr string
}

func (app *application) run() error {

	mux := http.NewServeMux()
	routes.Register(mux)

	err := http.ListenAndServe(app.conf.addr, mux)

	if err != nil {
		return err
	}
	return nil
}
