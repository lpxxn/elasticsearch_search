package es_lib

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/stretchr/testify/assert"
)

func TestDel1(t *testing.T) {
	body := struct {
		Name string `json:"name"`
		Age  int32  `json:"age"`
		Sort int32  `json:"sort"`
	}{}
	body.Name = "test zhang"
	body.Age = 18
	body.Sort = 123
	jsonBytes, _ := json.Marshal(body)
	documentID := "112"
	index := "del_test1"
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
	body.Sort = 16
	jsonBytes, _ = json.Marshal(body)
	req = esapi.IndexRequest{
		Index:      index,
		DocumentID: documentID,
		Body:       bytes.NewReader(jsonBytes),
	}
	resp, err = req.Do(context.Background(), ES7ClientT.Transport)
	t.Log(resp.String())
	t.Log(err)

	queryParam := map[string]interface{}{"query": map[string]interface{}{"term": map[string]interface{}{"age": 18}}}
	resp, err = ES7ClientT.SearchInfo(context.Background(), index, "_doc", queryParam)
	assert.Nil(t, err)
	t.Log(resp.String())

	doc := "_doc"
	err = ES7ClientT.DeleteByQueryInfo(context.Background(), index, queryParam, ES7ClientT.DeleteByQuery.WithDocumentType(doc))
	// , ES7ClientT.DeleteByQuery.WithConflicts("proceed")
	assert.Nil(t, err)
	if err != nil {
		// 可能存在相当多的失败实体，这个过程如果想统计版本冲突的次数而不是返回原因，可以在在URL中设置conflicts=proceed或者在请求体中带上"conflicts": "proceed"
		err = ES7ClientT.DeleteByQueryInfo(context.Background(), index, queryParam, ES7ClientT.DeleteByQuery.WithDocumentType(doc), ES7ClientT.DeleteByQuery.WithConflicts("proceed"))
		assert.Nil(t, err)
	}

	// update 是有 retry
	//ES7ClientT.Update.WithRetryOnConflict(10)

	resp, err = ES7ClientT.SearchInfo(context.Background(), index, "_doc", queryParam)
	assert.Nil(t, err)
	t.Log(resp.String())
}

/*
curl "localhost:9200/del_test1/_search?pretty"
curl "localhost:9200/del_test1/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "match": {"age": 18} },
   "sort": [
     { "age": "asc" }
   ]
 }
'

curl "localhost:9200/del_test1/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "term": {"age": 18} }
 }
'

curl -X POST "localhost:9200/del_test1/_doc?pretty" -H 'Content-Type: application/json' -d'
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
