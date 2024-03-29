package es_lib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/lpxxn/elasticsearch_search/es_lib/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ES7ClientT entity.Es7ClientType
var myTest2Index = "mytest2"

func TestInfo(t *testing.T) {
	resp, err := entity.Es7Client.Info()
	assert.Nil(t, err)
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	t.Log(string(body))
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func TestIndex(t *testing.T) {
	body := struct {
		PID       string   `json:"pid"`
		Name      string   `json:"name"`
		Age       int32    `json:"age"`
		Remark    string   `json:"remark"`
		City      string   `json:"city"`
		Location  Location `json:"location"`
		Location2 Location `json:"location2"`
	}{}
	body.Name = "wang_cai"
	body.Age = 5
	body.Remark = "一二三四"
	body.PID = "7ec0e0e5-a4b0-46d7-af56-5b3eab477aea"
	crateDoc := func(id string) {
		jsonBytes, _ := json.Marshal(body)
		//idxReq := es7api.IndexRequest{
		//	Index:        "my_test_3",
		//	DocumentType: "users",
		//	// 因为id 是固定的，所以每次会覆盖
		//	DocumentID: id,
		//	Body:       bytes.NewReader(jsonBytes),
		//}
		//resp, err := idxReq.Do(context.Background(), ES7ClientT.Client)
		//require.Nil(t, err)
		resp, err := ES7ClientT.CreateIndexDocument(context.Background(), "my_test_3", "users", id, jsonBytes)
		require.Nil(t, err)
		t.Log(resp.String())
	}
	crateDoc("a")

	body.Name = "diandian"
	body.Age = 1
	body.Remark = "李三张四测试"
	body.PID = "a4i0e0e5-a4b0-46d7-af56-5b3eab477aea"
	body.City = "beijing"
	// 在mapping.md里 location我有加到mapping里 geo_point，而location2没有就默认是 long 类型json
	body.Location = Location{
		Lat: 39.97621,
		Lon: 116.49176,
	}
	body.Location2 = Location{
		Lat: 39.97621,
		Lon: 116.49176,
	}
	crateDoc("b")

	body.Name = "xiaohei"
	body.Age = 3
	body.Remark = "a b c"
	body.PID = "9z80e0e5-b9b0-46d7-af56-5b3eab47icea"
	body.City = "beijing1"
	body.Location = Location{
		Lat: 36.97621,
		Lon: 106.49176,
	}
	body.Location2 = Location{
		Lat: 31.97621,
		Lon: 112.49176,
	}
	crateDoc("c")

	body.Name = "huahua"
	body.Age = 5
	body.Remark = "这个狗一的名字叫花花"
	body.PID = "qz8ie0f5-b9b0-9h0a-ad56-5i3eab47zcxa"
	body.City = "beijing2"
	body.Location = Location{
		Lat: 39.97621,
		Lon: 116.49176,
	}
	body.Location2 = Location{
		Lat: 39.97621,
		Lon: 116.49176,
	}
	crateDoc("d")

	body.Name = "dahuang"
	body.Age = 1
	body.Remark = "这个狗的一名字叫大黄"
	body.PID = "qz8ie0f5-b9k0-9h0a-ad06-5i3eab4izcxa"
	body.City = "beijing3"
	body.Location = Location{
		Lat: 29.97621,
		Lon: 116.49176,
	}
	body.Location2 = Location{
		Lat: 29.97621,
		Lon: 116.49176,
	}
	crateDoc("e")

	body.Name = "xiaobai"
	body.Age = 5
	body.Remark = "这个狗一的名字叫小白"
	body.PID = "qz8ie0f5-b0kz-ah0a-0906-5i3eab4izcxa"
	crateDoc("f")

	body.Name = "dingdang"
	body.Age = 15
	body.Remark = "这个狗一的名字叫叮当"
	body.PID = "qz8ie0f5-b0kz-ah0a-p9h6-5i3eab4izcxa"
	crateDoc("g")

	// 所有的数据
	// curl "localhost:9200/my_test_3/_search?pretty"
	// 所有 users的数据
	// curl "localhost:9200/my_test_3/users/_search?pretty"

	// user id 为 a 的数据
	// curl "localhost:9200/my_test_3/users/a?pretty"
	// 所有 age = 1的
	//curl -X GET "localhost:9200/my_test_3/_search?q=age:1&pretty"

	/*
		curl -X GET "localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d'
		{
		    "query" : {
		        "term" : { "age" : 5 }
		    }
		}
		'

	*/

	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"age": 5,
	//		},
	//	},
	//}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": map[string]interface{}{
					"match": map[string]interface{}{
						"age": 5,
					},
				},
			},
		},
	}
	//resp, err := ES7ClientT.Search(
	//	ES7ClientT.Search.WithIndex("my_test_3"),
	//	//ES7ClientT.Search.WithDocumentType("users"),
	//	ES7ClientT.Search.WithBody(es7util.NewJSONReader(query)),
	//	ES7ClientT.Search.WithPretty(),
	//	ES7ClientT.Search.WithSize(1),
	//	)
	time.Sleep(time.Second)
	resp, err := ES7ClientT.SearchInfo(context.Background(), "my_test_3", "users", query)
	assert.Nil(t, err)
	t.Log(resp)

	deleteQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"age": 5,
			},
		},
	}
	err = ES7ClientT.DeleteByQueryInfo(context.Background(), "my_test_3", deleteQuery, ES7ClientT.DeleteByQuery.WithConflicts("proceed"))
	assert.Nil(t, err)

	time.Sleep(time.Second)
	resp, err = ES7ClientT.SearchInfo(context.Background(), "my_test_3", "users", query)
	assert.Nil(t, err)
	t.Log(resp)

}

func TestIndex2(t *testing.T) {
	body := struct {
		Name string `json:"name"`
		Age  int32  `json:"age"`
	}{}
	body.Name = "wang_cai"
	body.Age = 2
	jsonBytes, _ := json.Marshal(body)
	// 这样的话 type 是_doc
	resp, err := ES7ClientT.Index("mytest2", bytes.NewReader(jsonBytes))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.String())

	resp, err = ES7ClientT.Index("mytest2", bytes.NewReader(jsonBytes))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.String())

	/*
			"hits" : [
			      {
			        "_index" : "mytest2",
			        "_type" : "_doc",
			        "_id" : "hgSJdXABXv3MPkS7vw-X",
			        "_score" : 1.0,
			        "_source" : {
			          "name" : "wang_cai",
			          "age" : 2
			        }
			      }
			    ]
			  }
		Elasticsearch 6: Rejecting mapping update as the final mapping would have more than 1 type

		Prior to elasticsearch v6, an index can have only 1 mapping by default.
		In previous version 5.x, multiple mapping was possible for an index.
		Though you may change this default setting by updating index setting
		"index.mapping.single_type": false
	*/
	/*
		curl -X GET "localhost:9200/mytest2/_search?_source=false&pretty"
		用 curl -X GET "localhost:9200/mytest2/_search?pretty" 得到所有数据，用— _id
		curl -X GET "localhost:9200/mytest2/_doc/VeVz_nUBqYxFvRaV80GB?&pretty"
		{
		  "_index" : "mytest2",
		  "_type" : "_doc",
		  "_id" : "VeVz_nUBqYxFvRaV80GB",
		  "_version" : 1,
		  "_seq_no" : 1,
		  "_primary_term" : 1,
		  "found" : true,
		  "_source" : {
		    "name" : "wang_cai",
		    "age" : 2
		  }
		}
	*/
}

// 在同一个id下多次创建，第一次返回的是 StatusCode 201 Created 再次index 返回200 OK
// err 都是 nil  curl -X GET "localhost:9200/mytest2/_doc/11?pretty"
func TestIndexDuplicated(t *testing.T) {
	body := struct {
		Name string `json:"name"`
		Age  int32  `json:"age"`
	}{}
	body.Name = "laotie"
	body.Age = 18
	jsonBytes, _ := json.Marshal(body)
	documentID := "112"
	index := "mytest2"
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: documentID,
		Body:       bytes.NewReader(jsonBytes),
	}
	resp, err := req.Do(context.Background(), ES7ClientT.Transport)
	t.Log(resp.String())
	t.Log(err)

	body.Name = "laotie2"
	body.Age = 18
	jsonBytes, _ = json.Marshal(body)
	req = esapi.IndexRequest{
		Index:      index,
		DocumentID: documentID,
		Body:       bytes.NewReader(jsonBytes),
	}
	resp, err = req.Do(context.Background(), ES7ClientT.Transport)
	t.Log(resp.String())
	t.Log(err)
}

func TestMain(m *testing.M) {
	ES7ClientT = entity.NewEsClient()

	os.Exit(m.Run())
}
