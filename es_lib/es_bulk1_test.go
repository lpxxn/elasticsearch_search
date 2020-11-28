package es_lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"testing"
)

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
		ID:   1,
		Name: "zhang",
		Age:  2,
	}, &body{
		ID:   3,
		Name: "wang",
		Age:  3,
	})
	count := 1000
	batch := 255
	var buf bytes.Buffer
	var numBatches int
	var currBatch int
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
		if i > 0 && i%batch == 0 || i == count-1 {
			doBulk(currBatch, numBatches, buf, indexName, documentType)
		}
	}
	if buf.Len() > 0 {
		doBulk(currBatch, numBatches, buf, indexName, documentType)
	}
	fmt.Println(buf.Len())
	fmt.Print("\n")
	log.Println(strings.Repeat("▔", 65))

}

func doBulk(currBatch int, numBatches int, buf bytes.Buffer, indexName string, documentType string) {
	fmt.Printf("[%d/%d] ", currBatch, numBatches)

	res, err := ES7ClientT.Bulk(bytes.NewReader(buf.Bytes()), ES7ClientT.Bulk.WithIndex(indexName), ES7ClientT.Bulk.WithDocumentType(documentType))
	if err != nil {
		log.Fatalf("Failure indexing batch %d: %s", currBatch, err)
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

 */
