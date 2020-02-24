package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/lpxxn/elasticsearch_search/es_lib"
)

func main() {
	es7 := es_lib.NewEsClient()
	resp, err := es_lib.Es7Client.Info()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.String())

	body := struct {
		Name string `json:"name"`
		Age int32	`json:"age"`
	}{}
	body.Name = "wang_cai"
	body.Age = 2
	jsonBytes, _ := json.Marshal(body)
	resp, err = es7.Index("mytest2", bytes.NewReader(jsonBytes))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.String())
}
