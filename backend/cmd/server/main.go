package main

import "github.com/DivyanshuShekhar55/yellow-monkey/backend/internal/es"

func main() {

	es := es.ConnectES()

	conf := config{
		addr: ":6969",
		es: es,
	}

	app := application{
		conf: conf,
	}

	app.run()

}
