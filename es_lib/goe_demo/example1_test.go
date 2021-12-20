package goe_demo

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	es7util "github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/lpxxn/elasticsearch_search/es_lib/entity"
)

var ES7ClientT entity.Es7ClientType

//var myTest2Index = "mytest_geo1"

func TestBoundingBoxQuery1(t *testing.T) {
	queryBody := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match_all": struct{}{},
				},
				"filter": map[string]interface{}{
					"geo_bounding_box": map[string]interface{}{
						"hotel.location": map[string]interface{}{
							"top_left": map[string]interface{}{
								"lat": 40.73,
								"lon": -74.1,
							},
							"bottom_right": map[string]interface{}{
								"lat": 40.01,
								"lon": -71.12,
							},
						},
					},
				},
			},
		},
	}
	resp, err := ES7ClientT.Search(ES7ClientT.Search.WithIndex("mytest_geo1"),
		ES7ClientT.Search.WithContext(context.Background()),
		ES7ClientT.Search.WithSize(200), ES7ClientT.Search.WithPretty(),
		ES7ClientT.Search.WithBody(es7util.NewJSONReader(queryBody)),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())

	queryBody = map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match_all": struct{}{},
				},
				"filter": map[string]interface{}{
					"geo_bounding_box": map[string]interface{}{
						"hotel.location": map[string]interface{}{
							"top_left": map[string]interface{}{
								"lat": 40.73,
								"lon": -74.1,
							},
							"bottom_right": map[string]interface{}{
								"lat": 40.01,
								"lon": -71.12,
							},
						},
					},
				},
			},
		},
	}

	resp, err = ES7ClientT.Search(ES7ClientT.Search.WithIndex("mytest_geo1,my_geoshapes"),
		ES7ClientT.Search.WithContext(context.Background()),
		ES7ClientT.Search.WithSize(200), ES7ClientT.Search.WithPretty(),
		ES7ClientT.Search.WithBody(es7util.NewJSONReader(queryBody)),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())
	b, _ := json.Marshal(queryBody)
	t.Log(string(b))
}

func TestLocations1(t *testing.T) {
	queryBody := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match_all": struct{}{},
				},
				"filter": map[string]interface{}{
					"geo_distance": map[string]interface{}{
						"distance":       "200km",
						"hotel.location": []int64{-70, 40},
					},
				},
			},
		},
	}
	resp, err := ES7ClientT.Search(ES7ClientT.Search.WithIndex("mytest_geo1,my_geoshapes"),
		ES7ClientT.Search.WithContext(context.Background()),
		ES7ClientT.Search.WithSize(200), ES7ClientT.Search.WithPretty(),
		ES7ClientT.Search.WithBody(es7util.NewJSONReader(queryBody)),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())
	b, _ := json.Marshal(queryBody)
	t.Log(string(b))
}

func TestMain(m *testing.M) {
	ES7ClientT = entity.NewEsClient()

	os.Exit(m.Run())
}
