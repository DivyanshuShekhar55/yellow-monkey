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
	Name     string   `json:"name"`
	Tags     []string `json:"tags"` // maybe enums??
	Location Coords   `json:"location"`
}

type GroupImpl struct {
	Conn *elasticsearch.Client
}

type GroupSearchResponse struct {
	Hits   int
	Values []Group
}

type SearchGroupRequestBody struct {
	Name     *string  `json:"name,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Location Coords   `json:"location"`
	MinRad   int      `json:"min_radius"`
	MaxRad   int      `json:"max_radius"`
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

func InsertGroup(ctx context.Context, conn *elasticsearch.Client, group *Group, refreshStrategy string) error {
	doc, err := json.Marshal(group)
	if err != nil {
		log.Printf("error parsing group %s", err)
		return fmt.Errorf("error parsing group %s", err)
	}

	req := esapi.IndexRequest{
		Index:   "groups",
		Body:    strings.NewReader(string(doc)),
		Refresh: refreshStrategy, // "false", "should_wait"
	}

	res, err := req.Do(ctx, conn)
	if err != nil {
		log.Printf("Error indexing group doc %s", err)
		return fmt.Errorf("error indexing group doc %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("error indexing group doc %s", res.String())
		return fmt.Errorf("error indexing group doc %s", err)
	}

	log.Printf("group indexed succesfully")
	fmt.Printf("group indexed succesfully")

	return nil
}

func SearchGroupByLocation(
	ctx context.Context,
	conn *elasticsearch.Client,
	groupName *string, // optional hai => if non present pointer = nil
	location Coords,
	minRad int,
	maxRad int,
	tags []string,
) *GroupSearchResponse {

	// max radius pe 7 km ka limit
	if maxRad < 7 {
		maxRad = 7
	}

	filters := []interface{}{}
	mustNot := []interface{}{}
	should := []interface{}{}

	if len(tags) > 0 {
		filters = append(filters, map[string]interface{}{
			"terms": map[string]interface{}{
				"tags": tags,
			},
		})
	}

	filters = append(filters, map[string]interface{}{
		"geo_distance": map[string]interface{}{
			"distance": fmt.Sprintf("%dkm", maxRad),
			"location": map[string]float64{
				"lat": location.Lat,
				"lon": location.Lon,
			},
		},
	})

	// minRad >= 0 always
	if minRad > 0 {
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

	if groupName != nil && *groupName != "" {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"name": map[string]interface{}{
					"query": *groupName,
					"boost": 1.45,
				},
			},
		})
	}

	// Build the query object
	queryObj := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter":   filters,
				"must_not": mustNot,
				"should":   should,
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
		log.Printf("error parsing group-search response %s", err)
		return nil
	}

	vals := make([]Group, 0, len(r.Hits.Hits))
	for _, hit := range r.Hits.Hits {
		vals = append(vals, hit.Source)
	}

	return &GroupSearchResponse{
		Hits:   r.Hits.Total.Value,
		Values: vals,
	}
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
