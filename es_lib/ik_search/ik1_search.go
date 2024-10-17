package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/opensearch-project/opensearch-go/v2"
	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
)

func main() {
	// 创建 OpenSearch 客户端
	client, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{"https://your-opensearch-endpoint:443"},
		// 根据需要添加认证信息
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 定义索引名称
	indexName := "my_index"

	// 创建索引并设置 mapping
	mapping := `{
		"mappings": {
			"properties": {
				"content": {
					"type": "text",
					"analyzer": "ik_smart"
				}
			}
		}
	}`

	req := opensearchapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(mapping),
	}

	res, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("Error creating index: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error creating index: %s", res.String())
	}

	// 执行模糊查询
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"content": map[string]interface{}{
					"query":     "你要搜索的内容",
					"fuzziness": "AUTO",
				},
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// 执行搜索请求
	res, err = client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(indexName),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// 打印错误信息
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// 打印搜索结果
	fmt.Printf("搜索结果: %+v\n", r["hits"].(map[string]interface{})["hits"])
}

/*
对于包含中文字符、数字等的字段，AWS OpenSearch 提供了几种适合的分词器选择：

标准分词器（Standard Analyzer）：这是默认的分词器，对于英文和数字效果不错，但对中文的支持有限。

IK 分词器：专门为中文设计的分词器，支持智能分词和最大化分词。

自定义分词器：可以根据需求组合不同的字符过滤器、分词器和词元过滤器。

对于中文内容，建议使用 IK 分词器。如果 OpenSearch 中没有预装 IK 分词器，你需要安装相应的插件。

关于 mapping，虽然 OpenSearch 会自动为新的索引和字段创建 mapping，但为了更好地控制字段的索引和搜索行为，建议手动配置 mapping。

以下是使用 Golang 进行模糊查询的示例，包括设置 mapping 和执行查询：
*/
