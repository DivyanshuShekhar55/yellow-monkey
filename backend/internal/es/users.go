package es

import (
	"context"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
)

type Coords struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type User struct {
	Username string `json:"username"`
	Location Coords `json:"location"`
}

type UserImpl struct {
	Conn *elasticsearch.Client
}

func (u *UserImpl) CreateUserIndex() {

	index := "users"

	// check if already exists
	req := esapi.IndicesExistsRequest{
		Index: []string{index},
	}

	res, err := req.Do(context.Background(), u.Conn)

	if err != nil {
		log.Printf("error checking existence of user index %s", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 404 {
		// index does not exists
		mapping := `{
			"mappings":{
			
				"properties" : {
					"username" : {
						"type" : "text"
					},
					"location" : {
						"type" : "geo_point"
					}
				}
			}
		}`

		req := esapi.IndicesCreateRequest{
			Index: index,
			Body:  strings.NewReader(mapping),
		}

		res, err := req.Do(context.Background(), u.Conn)
		if err != nil {
			log.Printf("Error creating user index %v", err)
			return
		}

		defer res.Body.Close()

		if res.IsError() {
			log.Printf("Error creating user index %v", err)
			return
		}

		log.Printf("created user index")
	} else {
		log.Printf("User index already exists")
	}
}

func (u *UserImpl) PutUser() {

}
