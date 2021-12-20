package goe_demo

import (
	"context"
	"os"
	"testing"

	es7util "github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/lpxxn/elasticsearch_search/es_lib/entity"
)
var ES7ClientT entity.Es7ClientType
var myTest2Index = "mytest_geo1"

func TestLocations1(t *testing.T) {
	resp, err := ES7ClientT.Search(ES7ClientT.Search.WithIndex("mytest_geo1,my_geoshapes"),
		ES7ClientT.Search.WithContext(context.Background()),
		ES7ClientT.Search.WithSize(200), ES7ClientT.Search.WithPretty(),
		ES7ClientT.Search.WithBody(es7util.NewJSONReader(map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": map[string]interface{}{
						"match_all": struct {}{},
					},
					"filter": map[string]interface{}{
						"geo_distance": map[string]interface{}{
							"distance": "200km",
							"hotel.location": []int64{-70, 40},
						},
					},
				},
			},

		})),
		)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())

}

func TestMain(m *testing.M) {
	ES7ClientT = entity.NewEsClient()

	os.Exit(m.Run())
}
