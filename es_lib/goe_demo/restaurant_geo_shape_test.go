package goe_demo

import (
	"context"
	"encoding/json"
	"testing"

	es7util "github.com/elastic/go-elasticsearch/v7/esutil"
)

func TestEnvelope1(t *testing.T) {
	queryBody := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match_all": struct{}{},
				},
				"filter": map[string]interface{}{
					"geo_shape": map[string]interface{}{
						"hotel.location": map[string]interface{}{
							"shape": map[string]interface{}{
								"type": "envelope",
								"coordinates": [][]float64{[]float64{116.49036,39.97665}, {116.49211,39.97584}},
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
	b, _ := json.Marshal(queryBody)
	t.Log(string(b))
}


func TestPlygonedit1(t *testing.T) {
	queryBody := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match_all": struct{}{},
				},
				"filter": map[string]interface{}{
					"geo_shape": map[string]interface{}{
						"hotel.location": map[string]interface{}{
							"relation": "within",
							"shape": map[string]interface{}{
								// "type": "polygon",
								"coordinates": [][][]float64{[][]float64{
									[]float64{116.49114,39.97654},
									[]float64{116.49118,39.97621},
									[]float64{116.49118,39.97585},
									[]float64{116.49208,39.97659},
									[]float64{116.49114,39.97654},
								}},
							},
						},
					},
				},
			},
		},
	}
	// []float64{
	//  北°, 东

	resp, err := ES7ClientT.Search(ES7ClientT.Search.WithIndex("mytest_geo1"),
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