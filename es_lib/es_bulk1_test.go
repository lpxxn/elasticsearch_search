package es_lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"testing"
)
// curl -X DELETE "localhost:9200/bulk_index_1?pretty"

func TestBulkIndex1(t *testing.T) {
	indexName := "bulk_index_1"
	documentType := "data"
	type body struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Age  int32  `json:"age"`
	}
	bodyArray := []*body{}

	bodyArray = append(bodyArray, &body{
		ID:   0,
		Name: "li",
		Age:  1,
	}, &body{
		ID:   2,
		Name: "zhang",
		Age:  2,
	}, &body{
		ID:   3,
		Name: "wang",
		Age:  3,
	}, &body{
		ID:   4,
		Name: "zhao",
		Age:  4,
	})
	batch := 255
	buf := new(bytes.Buffer)
	for i, v := range bodyArray {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%d" } }%s`, v.ID, "\n"))
		data, err := json.Marshal(v)
		if err != nil {
			log.Fatalf("Cannot encode article %d: %s", v.ID, err)
		}
		data = append(data, "\n"...) // <-- Comment out to trigger failure for batch
		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)
		// When a threshold is reached, execute the Bulk() request with body from buffer
		if i > 0 && i%batch == 0 {
			doBulk(buf, indexName, documentType)
		}
	}
	if buf.Len() > 0 {
		doBulk(buf, indexName, documentType)
	}
	fmt.Println(buf.Len())
	fmt.Println()
	log.Println(strings.Repeat("▔", 65))
}

func doBulk(buf *bytes.Buffer, indexName string, documentType string) {
	fmt.Printf("bytes: %s \n", buf.String())
	res, err := ES7ClientT.Bulk(bytes.NewReader(buf.Bytes()), ES7ClientT.Bulk.WithIndex(indexName), ES7ClientT.Bulk.WithDocumentType(documentType))
	if err != nil {
		log.Fatalf("Failure indexing batch  %#v", err)
	}
	if res.IsError() {
		var raw map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
			log.Fatalf("Failure to to parse response body: %s", err)
		} else {
			log.Printf("  Error: [%d] %s: %s",
				res.StatusCode,
				raw["error"].(map[string]interface{})["type"],
				raw["error"].(map[string]interface{})["reason"],
			)
		}
	} else {
		var blk *bulkResponse
		if err := json.NewDecoder(res.Body).Decode(&blk); err != nil {
			log.Fatalf("Failure to to parse response body: %s", err)
		} else {
			for _, d := range blk.Items {
				if d.Index.Status > 201 {
					// ... increment the error counter ...
					// ... and print the response status and error information ...
					log.Printf("  Error: [%d]: %s: %s: %s: %s",
						d.Index.Status,
						d.Index.Error.Type,
						d.Index.Error.Reason,
						d.Index.Error.Cause.Type,
						d.Index.Error.Cause.Reason,
					)
				}
			}
		}
	}

	// Close the response body, to prevent reaching the limit for goroutines or file handles
	res.Body.Close()

	// Reset the buffer and items counter
	buf.Reset()
}


type bulkResponse struct {
	Errors bool `json:"errors"`
	Items  []struct {
		Index struct {
			ID     string `json:"_id"`
			Result string `json:"result"`
			Status int    `json:"status"`
			Error  struct {
				Type   string `json:"type"`
				Reason string `json:"reason"`
				Cause  struct {
					Type   string `json:"type"`
					Reason string `json:"reason"`
				} `json:"caused_by"`
			} `json:"error"`
		} `json:"index"`
	} `json:"items"`
}

/*
https://golangbyexample.com/go-iterator-design-pattern/
https://www.tutorialspoint.com/design_pattern/iterator_pattern.htm
 */
