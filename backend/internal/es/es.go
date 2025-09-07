package es

import (
	"log"

	"github.com/elastic/go-elasticsearch/v9"
)

type ESClient struct {
	Conn *elasticsearch.Client
	Users UserFn
}

func NewESClient() *ESClient {
	conn := ConnectES()
	if conn == nil { return nil }

	return &ESClient{
		Conn: conn,
		Users: &UserImpl{Conn: conn},
	}
}

type UserFn interface {
	CreateUserIndex()
	PutUser()
}

func ConnectES() *elasticsearch.Client{
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		// ideally kuch aur bhi karenge yaha like retry
		log.Fatal("couldn't connect to elasticsearch")
		return nil
	}

	res, err:= es.Info()
	if err!=nil{
		// retry or other logic
		log.Fatal("couldn't get info from elastic")
		return nil
	}
	defer res.Body.Close()

	return es

}

