package es_lib

import (
	"context"
	"errors"
	"fmt"
	"sync"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	es7api "github.com/elastic/go-elasticsearch/v7/esapi"
)

type es7Client struct {
	*elasticsearch7.Client
}

var Es7Client *es7Client
var es7connOnce sync.Once

func NewEsClient() *elasticsearch7.Client {
	if Es7Client != nil {
		return Es7Client.Client
	}
	es7connOnce.Do(func() {
		// elasticsearch7.NewDefaultClient()
		cfg := elasticsearch7.Config{
			Addresses: []string{
				"http://localhost:9200",
				//"http://localhost:9201",
			},
			Username: "elastic",
			Password: "changeme",
			// ...
		}
		es7, err := elasticsearch7.NewClient(cfg)
		if err != nil {
			panic(err)
		}
		Es7Client = &es7Client{Client: es7}
	})
	info, err := Es7Client.Info()
	if err != nil {
		panic(err)
	}
	fmt.Println(info.String())
	return Es7Client.Client
}

func (e *es7Client) Info() (*es7api.Response, error) {
	req := es7api.InfoRequest{
		Pretty:     true,
		Human:      true,
		ErrorTrace: false,
	}
	resp, err := req.Do(context.Background(), e.Client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.IsError() {
		return nil, errors.New(resp.String())
	}
	return resp, nil
}

func (e *es7Client) SearchInfo() (*es7api.Response, error) {

	return nil, nil
}

func (e *es7Client) CreateIndexDocument(index, body string) error {
	//req := e.Index()
	return nil
}