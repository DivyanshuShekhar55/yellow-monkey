package main

import (
	"context"
	"log"
	"net/http"

	"github.com/DivyanshuShekhar55/yellow-monkey/backend/internal/db"
	esmodule "github.com/DivyanshuShekhar55/yellow-monkey/backend/internal/es"
	"github.com/DivyanshuShekhar55/yellow-monkey/backend/internal/routes"
)

func main() {

	es := esmodule.ConnectES()
	ctx := context.Background()
	pool := db.ConnectPG(ctx)

	handler := routes.NewHandler(es, pool)

	mux := http.NewServeMux()
	handler.Register(mux)

	conf := config{
		addr:   ":6969",
		es:     es,
		pgpool: pool,
	}

	app := application{
		conf: conf,
	}

	esmodule.CreateGroupIndex(context.Background(), es)
	esmodule.CreateUserIndex(es)

	err := http.ListenAndServe(app.conf.addr, mux)

	if err != nil {
		log.Fatalf("err starting %v", err)
	}

}
