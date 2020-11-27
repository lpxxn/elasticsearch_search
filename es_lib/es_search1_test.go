package es_lib

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/stretchr/testify/assert"
)

func TestSearch1(t *testing.T) {
	body := struct {
		Name string `json:"name"`
		Age  int32  `json:"age"`
		Sort int32  `json:"sort"`
	}{}
	body.Name = "test zhang"
	body.Age = 18
	body.Sort = 12
	jsonBytes, _ := json.Marshal(body)
	documentID := "112"
	index := "search_test1"
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: "1",
		Body:       bytes.NewReader(jsonBytes),
	}
	resp, err := req.Do(context.Background(), ES7ClientT.Transport)
	t.Log(resp.String())
	t.Log(err)

	body.Name = "laotie 2 li"
	body.Age = 18
	body.Sort = 2
	jsonBytes, _ = json.Marshal(body)
	req = esapi.IndexRequest{
		Index:      index,
		DocumentID: "2",
		Body:       bytes.NewReader(jsonBytes),
	}
	resp, err = req.Do(context.Background(), ES7ClientT.Transport)
	t.Log(resp.String())
	t.Log(err)

	body.Name = "laotie 3 wang"
	body.Age = 188
	body.Sort = 7
	jsonBytes, _ = json.Marshal(body)
	req = esapi.IndexRequest{
		Index:      index,
		DocumentID: documentID,
		Body:       bytes.NewReader(jsonBytes),
	}
	resp, err = req.Do(context.Background(), ES7ClientT.Transport)
	t.Log(resp.String())
	t.Log(err)

	resp, err = ES7ClientT.SearchInfo(context.Background(), index, "_doc",
		map[string]interface{}{"query": map[string]interface{}{"match": map[string]interface{}{"name": "laotie, zhang"}}}) // 只要有 laotie 或者 zhang 就会查出来
	assert.Nil(t, err)
	t.Log(resp.String())

	resp, err = ES7ClientT.SearchInfo(context.Background(), index, "_doc",
		map[string]interface{}{"query": map[string]interface{}{"match": map[string]interface{}{"name": "laotie zhang"}}}) // 只要有 laotie 或者 zhang 就会查出来和上面是一样的
	assert.Nil(t, err)
	t.Log(resp.String())

	resp, err = ES7ClientT.SearchInfo(context.Background(), index, "_doc",
		map[string]interface{}{"query": map[string]interface{}{"constant_score": map[string]interface{}{
			"filter": map[string]interface{}{"match": map[string]interface{}{"name": "laotie zhang"}}}}}) // 只要有 laotie 或者 zhang 就会查出来和上面是一样的
	assert.Nil(t, err)
	t.Log(resp.String())

	resp, err = ES7ClientT.SearchInfo(context.Background(), index, "_doc",
		map[string]interface{}{"query": map[string]interface{}{"term": map[string]interface{}{"name": "wang"}}}) // term是只要有 wang
	assert.Nil(t, err)
	t.Log(resp.String())

}

/*
curl "localhost:9200/search_test1/_search?pretty"
curl "localhost:9200/search_test1/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "match": {"age": 18} },
   "sort": [
     { "age": "asc" }
   ]
 }
'

curl "localhost:9200/search_test1/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": {
		"constant_score": {
	        "filter" : {
				"match": {"age": 18}
			}
		}
	},
   "sort": [
        { "_id": "asc" }
   ]
 }
'

curl "localhost:9200/search_test1/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": {
		"constant_score": {
	        "filter" : {
				"match": {"age": 18}
			}
		}
	},
   "sort": [
        { "sort" : { "order": "asc", "unmapped_type" : "integer"} }
   ]
 }
'
// No mapping found for [sort] in order to sort on 要加上 unmapped_type

curl "localhost:9200/search_test1/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "term": {"age": 18} }
 }
'

curl -X POST "localhost:9200/search_test1/_doc?pretty" -H 'Content-Type: application/json' -d'
{
  "query": {
    "match": {
      "user.id": "elkbee"
    }
  }
}
'

curl -X POST "localhost:9200/my-index-000001/_delete_by_query?pretty" -H 'Content-Type: application/json' -d'
{
  "query": {
    "match": {
      "user.id": "elkbee"
    }
  }
}
'

*/
