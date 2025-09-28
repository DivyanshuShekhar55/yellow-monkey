package main

import (
	"context"
	"log"
	"net/http"

	esmodule "github.com/DivyanshuShekhar55/yellow-monkey/backend/internal/es"
	"github.com/DivyanshuShekhar55/yellow-monkey/backend/internal/routes"
)

func main() {

	es := esmodule.ConnectES()
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

	esmodule.CreateGroupIndex(context.Background(), es)

	err := http.ListenAndServe(app.conf.addr, mux)

	if err != nil {
		log.Fatalf("err starting %v", err)
	}

}
