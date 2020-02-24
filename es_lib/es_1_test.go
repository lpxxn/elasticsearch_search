package es_lib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"
	es7api "github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/stretchr/testify/require"
)

var es7 *elasticsearch.Client
var myTest2Index = "mytest2"
func TestIndex(t *testing.T) {
	body := struct {
		PID string `json:"pid"`
		Name string `json:"name"`
		Age  int32  `json:"age"`
		Remark string `json:"remark"`
	}{}
	body.Name = "wang_cai"
	body.Age = 5
	body.Remark = "一二三四"
	body.PID = "7ec0e0e5-a4b0-46d7-af56-5b3eab477aea"
	crateDoc := func(id string) {
		jsonBytes, _ := json.Marshal(body)
		idxReq := es7api.IndexRequest{
			Index:               "my_test_3",
			DocumentType:        "users",
			// 因为id 是固定的，所以每次会覆盖
			DocumentID:          id,
			Body:                bytes.NewReader(jsonBytes),
		}
		resp, err := idxReq.Do(context.Background(), es7)
		require.Nil(t, err)
		t.Log(resp.String())
	}
	crateDoc("a")

	body.Name = "diandian"
	body.Age = 1
	body.Remark = "李三张四测试"
	body.PID = "a4i0e0e5-a4b0-46d7-af56-5b3eab477aea"
	crateDoc("b")

	body.Name = "xiaohei"
	body.Age = 3
	body.Remark = "a b c"
	body.PID = "9z80e0e5-b9b0-46d7-af56-5b3eab47icea"
	crateDoc("c")

	body.Name = "huahua"
	body.Age = 5
	body.Remark = "这个狗的名字叫花花"
	body.PID = "qz8ie0f5-b9b0-9h0a-ad56-5i3eab47zcxa"
	crateDoc("d")


	// 所有的数据
	// curl "localhost:9200/my_test_3/_search?pretty"
	// 所有 users的数据
	// curl "localhost:9200/my_test_3/users/_search?pretty"

	// user id 为 a 的数据
	// curl "localhost:9200/my_test_3/users/a?pretty"
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
	resp, err := es7.Index("mytest2", bytes.NewReader(jsonBytes))
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
}

func TestMain(m *testing.M) {
	es7 = NewEsClient()

	os.Exit(m.Run())
}
