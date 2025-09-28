package es

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
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

type GroupSearchResponse struct {
	Hits   int
	Values []Group
}

func CreateGroupIndex(ctx context.Context, conn *elasticsearch.Client) {
	index := "groups"

	//check for already existing index:
	req := esapi.IndicesExistsRequest{
		Index: []string{index},
	}

	res, err := req.Do(ctx, conn)
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
					"type": "keyword"
				}
			}
		}
	}`

		req := esapi.IndicesCreateRequest{
			Index: index,
			Body:  strings.NewReader(mapping),
		}

		res, err := req.Do(ctx, conn)
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

func InsertGroup(ctx context.Context, conn *elasticsearch.Client, group *Group, refreshStrategy string) {
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

	res, err := req.Do(ctx, conn)
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

func SearchGroupByLocation(ctx context.Context, conn *elasticsearch.Client ,groupName string, location Coords, minRad, maxRad int, tag ...string) *GroupSearchResponse {
	// [tags] + [use geohash to filter fast closeby areas only] + [user loc + start dis + end dis (with max limit)] + [name]

	// Prepare tags
	tags := []string{"tag1", "tag2"}
	tagsJSON, _ := json.Marshal(tags) // becomes: ["tag1","tag2"]

	maxRad = max(maxRad, 7)

	// WARNING : GANDMASTI AHEAD
	queryObj := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []interface{}{
					map[string]interface{}{"terms": map[string]interface{}{"tags": tagsJSON}},
					map[string]interface{}{
						"location":   map[string]float64{"lat": location.Lat, "lon": location.Lon},
						"neighbours": true,
						"precision":  "2km", // level 6 precision, saves space
					},
					map[string]interface{}{
						"geo_distance_range": map[string]interface{}{
							"gte":      strconv.Itoa(minRad) + "km",
							"lte":      strconv.Itoa(maxRad) + "km",
							"location": map[string]float64{"lat": location.Lat, "lon": location.Lon},
						},
					},
				},
				"should": []interface{}{
					map[string]interface{}{
						"match": map[string]interface{}{
							"name": map[string]interface{}{
								"query": groupName,
								"boost": 1.45,
							},
						},
					},
				},
			},
		},
	}
	queryBytes, err := json.Marshal(queryObj)
	if err != nil {
		log.Printf("error marshalling group search-query %s", err)
		return nil
	}

	req := esapi.SearchRequest{
		Index:          []string{"groups"},
		Body:           strings.NewReader(string(queryBytes)),
		TrackTotalHits: "true",
	}

	res, err := req.Do(ctx, conn)
	if err != nil {
		log.Printf("cannot search for group %s", err)
		return nil
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Printf("cannot search for group %s", res.String())
		return nil
	}

	// parse the response, bhejne se pahle
	var r struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source Group `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("error parsing group-serach response %s", err)
		return nil
	}

	var vals []Group
	for _, hit := range r.Hits.Hits {
		vals = append(vals, hit.Source)
	}

	search_res := &GroupSearchResponse{
		Hits:   r.Hits.Total.Value,
		Values: vals,
	}

	return search_res

}

func GetAllGroups(ctx context.Context, conn *elasticsearch.Client) (*GroupSearchResponse, error) {

	query := `{
		"query": {
			"match_all": {}
		}
	}`
	req := esapi.SearchRequest{
		Index:          []string{"groups"},
		Body:           strings.NewReader(query),
		TrackTotalHits: "true",
	}

	res, err := req.Do(ctx, conn)
	if err != nil {
		return nil, fmt.Errorf("couldn't fetch groups %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error fetching data %s", res.String())
	}

	var r struct {
		Hits struct {
			Total struct {
				Value int `json:"total"`
			}
			Hits []struct {
				Source Group `json:"_source"`
			}
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	var vals []Group
	for _, hit := range r.Hits.Hits {
		vals = append(vals, hit.Source)
	}

	return &GroupSearchResponse{
		Hits:   r.Hits.Total.Value,
		Values: vals,
	}, nil

}
