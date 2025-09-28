package main

import (
	"log"
	"net/http"

	"github.com/DivyanshuShekhar55/yellow-monkey/backend/internal/es"
	"github.com/DivyanshuShekhar55/yellow-monkey/backend/internal/routes"
)

func main() {

	es := es.ConnectES()
	handler := routes.NewHandler(es)
	
	mux := http.NewServeMux()
	handler.Register(mux)


	conf := config{
		addr: ":6969",
		es:   es,
	}

	app := application{
		conf: conf,
	}

	err := http.ListenAndServe(app.conf.addr, mux)

	if err != nil {
		log.Fatalf("err starting %v", err)
	}

}
