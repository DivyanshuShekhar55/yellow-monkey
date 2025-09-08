package es

import (
	"context"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
)

type Group struct {
	Name     string
	Tags     []string // maybe enums??
	Location Coords
}

type GroupImpl struct {
	Conn *elasticsearch.Client
}

func (g *GroupImpl) CreateGroupIndex() {
	index := "groups"

	//check for already existing index:
	req := esapi.IndicesExistsRequest{
		Index: []string{index},
	}

	res, err := req.Do(context.Background(), g.Conn)
	if err != nil {
		log.Printf("couldn't search for existing group index %s", res)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 404 {
		// index doesnt exists create a new one

		mapping := `{
		"mappings": {
			"properties": {
				"name": {
					"type": "text"
				},
				"location": {
					"type": "geo_point"
				},
				"tags": {
					"type": []
				}
			}
		}
	}`

		req := esapi.IndicesCreateRequest{
			Index: index,
			Body:  strings.NewReader(mapping),
		}

		res, err := req.Do(context.Background(), g.Conn)
		if err != nil {
			log.Printf("Error creating user index %v", err)
			return
		}

		defer res.Body.Close()

		if res.IsError() {
			log.Printf("Error creating user index %v", res.String())
			return
		}

		log.Printf("created user index")
		log.Printf("user index succesfully created")

	} else {
		
	}

}
