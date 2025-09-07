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

type Coords struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type User struct {
	Username string `json:"username"`
	Location Coords `json:"location"`
}

type SearchUserResponse struct {
	Hits int64
	Values []User
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
			log.Printf("Error creating user index %v", res.String())
			return
		}

		log.Printf("created user index")
		log.Printf("user index succesfully created")

	} else {
		log.Printf("User index already exists")
	}
}

func (u *UserImpl) PutUser(user User) {
	doc, err := json.Marshal(user)
	if err != nil {
		log.Printf("error marshalling user schema %s", err)
		return
	}

	req := esapi.IndexRequest{
		Index:   "users",
		Body:    strings.NewReader(string(doc)),
		Refresh: "false",
	}

	res, err := req.Do(context.Background(), u.Conn)
	if err != nil {
		log.Printf("Error indexing user doc %s", err)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("error indexing user doc %s", res.String())
		return
	}

	log.Printf("user indexed succesfully")
	fmt.Printf("user indexed succesfully")

}

func (u *UserImpl) SearchUserByUsername(username string) *SearchUserResponse {
	query := fmt.Sprintf(`{
		"query": {
			"match": {
				"username": %s
			}
		}
	}`, username)

	req := esapi.SearchRequest{
		Index:          []string{"users"},
		Body:           strings.NewReader(query),
		TrackTotalHits: "true",
	}

	res, err := req.Do(context.Background(), u.Conn)
	if err != nil {
		log.Printf("error searching user doc %s", err)
		return nil
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("error searching user doc %s", res.String())
		return nil
	}

	// parse the response
	var r struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source User `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err!= nil{
		log.Printf("couldn't parse user response %s in search fn", err)
		return nil
	}

	var vals []User
	for _, hit := range r.Hits.Hits {
		vals = append(vals, hit.Source)
	}
	users := SearchUserResponse {
		Hits: int64(r.Hits.Total.Value),
		Values: vals,
	}

	return &users

}
