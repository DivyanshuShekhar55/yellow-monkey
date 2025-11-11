package main

import (
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/jackc/pgx/v5/pgxpool"
)

type application struct {
	conf config
}

type config struct {
	addr string
	//db   dbConfig
	es *elasticsearch.Client
	pgpool *pgxpool.Pool
}

type dbConfig struct {
	addr string
}

