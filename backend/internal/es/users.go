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
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
}

type SearchUserResponse struct {
	Hits   int
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
		// ideally define more options like num of shards and stuff
		mapping := `{
			"mappings":{
			
				"properties" : {
					"username" : {
						"type" : "text"
					},
					"location" : {
						"type" : "geo_point"
					},
					"age: {
						"type": "integer"
					},
					"gender": {
						"type": "keyword"
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

	// this is better than using sprintf
	query, _ := json.Marshal(map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]string{
				"username": username,
			},
		},
	})

	req := esapi.SearchRequest{
		Index:          []string{"users"},
		Body:           strings.NewReader(string(query)),
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

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("couldn't parse user response %s in search fn", err)
		return nil
	}

	var vals []User
	for _, hit := range r.Hits.Hits {
		vals = append(vals, hit.Source)
	}
	users := SearchUserResponse{
		Hits:   r.Hits.Total.Value,
		Values: vals,
	}

	return &users

}

func SearchUsersByLocation(ctx context.Context,
	conn *elasticsearch.Client,
	location Coords,
	minRad, maxRad int,
	minAge, maxAge int,
	gender string,
) *SearchUserResponse {

	// max radius ko clamp karna hai to 7kms
	if maxRad > 7 {
		maxRad = 7
	}
	if maxAge > 60 {
		maxAge = 60
	}

	filters := []interface{}{}
	//should := []interface{}{}
	mustNot := []interface{}{}

	filters = append(filters, map[string]interface{}{
		"term": map[string]interface{}{
			"gender": gender,
		},
	})

	filters = append(filters, map[string]interface{}{
		"range": map[string]interface{}{
			"age": map[string]interface{}{
				"gte": minAge,
				"lte": maxAge,
			},
		},
	})

	filters = append(filters, map[string]interface{}{
		"geo_distance": map[string]interface{}{
			"distance": fmt.Sprintf("%dkm", maxRad),
			"location": map[string]float64{
				"lat": location.Lat,
				"lon": location.Lon,
			},
		},
	})

	if minRad >= 0 {
		mustNot = append(mustNot, map[string]interface{}{
			"geo_distance": map[string]interface{}{
				"distance": fmt.Sprintf("%dkm", minRad),
				"location": map[string]float64{
					"lat": location.Lat,
					"lon": location.Lon,
				},
			},
		})
	}

	// build the query object
	queryObj := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter":   filters,
				"must_not": mustNot,
			},
		},
	}

	queryBytes, err := json.Marshal(queryObj)
	if err != nil {
		log.Printf("error marshalling user search req %v", err)
		return nil
	}

	req := esapi.SearchRequest{
		Index:          []string{"users"},
		Body:           strings.NewReader(string(queryBytes)),
		TrackTotalHits: "true",
	}

	res, err := req.Do(ctx, conn)
	if err != nil {
		log.Printf("cannot search for users %s", err)
		return nil
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("cannot search for users %s", res.String())
		return nil
	}

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

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("error parsing user-search response %s", err)
		return nil
	}

	vals := make([]User, 0, len(r.Hits.Hits))
	for _, hit := range r.Hits.Hits {
		vals = append(vals, hit.Source)
	}

	return &SearchUserResponse{
		Hits:   r.Hits.Total.Value,
		Values: vals,
	}

}
