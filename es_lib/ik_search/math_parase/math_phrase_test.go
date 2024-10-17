package math_parase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/opensearch-project/opensearch-go/v2"
	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
)

type Document struct {
	Content string `json:"content"`
}

func Test_MathPharase1(t *testing.T) {
	// 创建 OpenSearch 客户端
	cfg := opensearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	client, err := opensearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 准备索引和文档
	indexName := "test_index"
	prepareIndex(client, indexName)

	// 执行不同类型的查询并比较结果
	searchTerm := "OpenSearch 高级查询"

	fmt.Println("1. 使用 match 查询:")
	performSearch(client, indexName, buildMatchQuery(searchTerm))

	fmt.Println("\n2. 使用基本的 match_phrase 查询:")
	performSearch(client, indexName, buildMatchPhraseQuery(searchTerm, 0))

	fmt.Println("\n3. 使用带有 slop 的 match_phrase 查询:")
	performSearch(client, indexName, buildMatchPhraseQuery(searchTerm, 1))
}

func prepareIndex(client *opensearch.Client, indexName string) {
	// 删除已存在的索引
	deleteIndex(client, indexName)

	// 创建新索引，设置 ik_smart 分词器
	createIndex(client, indexName)

	// 添加示例文档
	addDocument(client, indexName, Document{Content: "OpenSearch 是一个强大的搜索引擎"})
	addDocument(client, indexName, Document{Content: "OpenSearch 提供高级查询功能"})
	addDocument(client, indexName, Document{Content: "使用 OpenSearch 进行高级查询和分析"})
	addDocument(client, indexName, Document{Content: "高级 OpenSearch 查询技巧"})
	addDocument(client, indexName, Document{Content: "高实在是高"})
	addDocument(client, indexName, Document{Content: "大学的高等数学"})
}

func deleteIndex(client *opensearch.Client, indexName string) {
	req := opensearchapi.IndicesDeleteRequest{
		Index: []string{indexName},
	}
	_, err := req.Do(context.Background(), client)
	if err != nil {
		log.Printf("Error deleting index: %s", err)
	}
}

func createIndex(client *opensearch.Client, indexName string) {
	// 定义索引映射，使用自定义分析器
	mapping := `{
        "settings": {
            "analysis": {
                "analyzer": {
                    "custom_analyzer": {
                        "type": "custom",
                        "tokenizer": "ik_smart",
                        "filter": ["lowercase"]
                    }
                }
            }
        },
        "mappings": {
            "properties": {
                "content": {
                    "type": "text",
                    "analyzer": "custom_analyzer",
                    "search_analyzer": "custom_analyzer"
                }
            }
        }
    }`

	req := opensearchapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(mapping),
	}
	_, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("Error creating index: %s", err)
	}
}

// 多个分词器
func createIndex2(client *opensearch.Client, indexName string) {
	// 定义索引映射，使用多字段
	mapping := `{
        "settings": {
            "analysis": {
                "analyzer": {
                    "ik_smart_lowercase": {
                        "type": "custom",
                        "tokenizer": "ik_smart",
                        "filter": ["lowercase"]
                    }
                }
            }
        },
        "mappings": {
            "properties": {
                "content": {
                    "type": "text",
                    "analyzer": "ik_smart_lowercase",
                    "fields": {
                        "standard": {
                            "type": "text",
                            "analyzer": "standard"
                        }
                    }
                }
            }
        }
    }`

	req := opensearchapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(mapping),
	}
	_, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("Error creating index: %s", err)
	}
}

func buildMultiFieldQuery(searchTerm string) map[string]interface{} {
	return map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  searchTerm,
				"fields": []string{"content", "content.standard"},
			},
		},
	}
}

func addDocument(client *opensearch.Client, indexName string, doc Document) {
	body, _ := json.Marshal(doc)
	req := opensearchapi.IndexRequest{
		Index: indexName,
		Body:  strings.NewReader(string(body)),
	}
	_, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("Error adding document: %s", err)
	}
}

func buildMatchQuery(searchTerm string) map[string]interface{} {
	return map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"content": searchTerm,
			},
		},
	}
}

func buildMatchPhraseQuery(searchTerm string, slop int) map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"content": map[string]interface{}{
					"query": searchTerm,
				},
			},
		},
	}

	if slop > 0 {
		query["query"].(map[string]interface{})["match_phrase"].(map[string]interface{})["content"].(map[string]interface{})["slop"] = slop
	}

	return query
}

func performSearch(client *opensearch.Client, indexName string, query map[string]interface{}) {
	body, _ := json.Marshal(query)
	req := opensearchapi.SearchRequest{
		Index: []string{indexName},
		Body:  strings.NewReader(string(body)),
	}

	res, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("Error performing search: %s", err)
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		fmt.Printf("- %s\n", source["content"])
	}
}
