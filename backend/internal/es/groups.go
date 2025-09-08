package es

import (
	"context"
	"encoding/json"
	"fmt"
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
			log.Printf("Error creating group index %v", err)
			return
		}

		defer res.Body.Close()

		if res.IsError() {
			log.Printf("Error creating group index %v", res.String())
			return
		}

		log.Printf("created group index")
		log.Printf("group index succesfully created")
		return

	} else {
		log.Printf("Group index already exists")
		return
	}

}

func (g *GroupImpl) InsertGroup(group *Group, refreshStrategy string) {
	doc, err := json.Marshal(group)
	if err != nil {
		log.Printf("error parsing group %s", err)
		return
	}

	req := esapi.IndexRequest{
		Index:   "groups",
		Body:    strings.NewReader(string(doc)),
		Refresh: refreshStrategy, // "false", "should_wait"
	}

	res, err := req.Do(context.Background(), g.Conn)
	if err != nil {
		log.Printf("Error indexing group doc %s", err)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("error indexing group doc %s", res.String())
		return
	}

	log.Printf("group indexed succesfully")
	fmt.Printf("group indexed succesfully")
}
