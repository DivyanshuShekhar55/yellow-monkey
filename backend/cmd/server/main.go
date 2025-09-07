package main

func main() {

	conf := config{
		addr: ":6969",
	}

	app := application{
		conf: conf,
	}

	app.run()

}
