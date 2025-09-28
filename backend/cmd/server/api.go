package main

import (
	"github.com/elastic/go-elasticsearch/v9"
)

type application struct {
	conf config
}

type config struct {
	addr string
	//db   dbConfig
	es *elasticsearch.Client
}

type dbConfig struct {
	addr string
}

