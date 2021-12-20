package entity

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	es7api "github.com/elastic/go-elasticsearch/v7/esapi"
	es7util "github.com/elastic/go-elasticsearch/v7/esutil"
)

type es7Client struct {
	*elasticsearch7.Client
}

type Es7ClientType = *es7Client

const esReqTimeout = time.Second * 5

var Es7Client Es7ClientType
var es7connOnce sync.Once

func NewEsClient() *es7Client {
	if Es7Client != nil {
		return Es7Client
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
	//info, err := Es7Client.Info()
	//defer info.Body.Close()
	//if err != nil {
	//	 panic(err)
	//}
	//fmt.Println(info.String())
	return Es7Client
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
	//defer resp.Body.Close()
	if resp.IsError() {
		return nil, errors.New(resp.String())
	}
	return resp, nil
}

func (e *es7Client) SearchInfo(ctx context.Context, index, docType string, query map[string]interface{}, o ...func(*es7api.SearchRequest)) (*es7api.Response, error) {
	opt := append(o, e.Search.WithIndex(index), e.Search.WithDocumentType(docType),
		e.Search.WithContext(ctx),
		e.Search.WithBody(es7util.NewJSONReader(query)),
		e.Search.WithSize(200), e.Search.WithPretty())
	resp, err := e.Search(opt...)
	return resp, err
}

func (e *es7Client) Batch(ctx context.Context) {
	//e.Bulk()
}

func (e *es7Client) CreateIndexDocument(ctx context.Context, index, documentType, documentID string, body []byte) (*es7api.Response, error) {
	req := es7api.IndexRequest{
		Index:        index,
		DocumentType: documentType,
		DocumentID:   documentID,
		Body:         bytes.NewReader(body),
		Timeout:      esReqTimeout,
	}
	resp, err := req.Do(ctx, e)
	if err != nil {
		return nil, err
	}
	if err = errByResponse(resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (e *es7Client) MGet(ctx context.Context) {
	//e.Client.Mget()
}

func (e *es7Client) DeleteByQueryInfo(ctx context.Context, index string, query map[string]interface{}, o ...func(*es7api.DeleteByQueryRequest)) error {
	opt := append(o,
		e.DeleteByQuery.WithContext(ctx),
		e.DeleteByQuery.WithHuman(),
	)
	resp, err := e.DeleteByQuery([]string{index}, es7util.NewJSONReader(query), opt...)
	if err != nil {
		return err
	}
	if err = errByResponse(resp); err != nil {
		return err
	}
	return nil
}

func errByResponse(resp *es7api.Response) error {
	if resp.IsError() {
		return errors.New(fmt.Sprintf("statusCode %s, error: %s", resp.Status(), resp.String()))
	}
	return nil
}
