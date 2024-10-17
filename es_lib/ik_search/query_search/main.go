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

const indexName = "test_chinese_index"

func main() {
	// 创建 OpenSearch 客户端
	client, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{"https://your-opensearch-endpoint:443"},
		// 根据需要添加认证信息
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 1. 创建索引
	if err := createIndex(client); err != nil {
		log.Fatalf("Error creating index: %s", err)
	}

	// 2. 插入测试数据
	if err := insertTestData(client); err != nil {
		log.Fatalf("Error inserting test data: %s", err)
	}

	// 3. 执行搜索
	if err := searchDocuments(client, "中文搜索测试"); err != nil {
		log.Fatalf("Error searching documents: %s", err)
	}
}

func createIndex(client *opensearch.Client) error {
	indexMapping := `{
		"settings": {
			"analysis": {
				"analyzer": {
					"ik_analyzer": {
						"type": "custom",
						"tokenizer": "ik_max_word",
						"filter": ["lowercase"]
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"title": {
					"type": "text",
					"analyzer": "ik_analyzer"
				},
				"content": {
					"type": "text",
					"analyzer": "ik_analyzer"
				}
			}
		}
	}`

	req := opensearchapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(indexMapping),
	}

	res, err := req.Do(context.Background(), client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error creating index: %s", res.String())
	}

	fmt.Println("Index created successfully")
	return nil
}

func insertTestData(client *opensearch.Client) error {
	documents := []map[string]interface{}{
		{
			"title":   "OpenSearch 中文测试",
			"content": "这是一个关于 OpenSearch 和中文搜索的测试文档。",
		},
		{
			"title":   "搜索引擎技术",
			"content": "现代搜索引擎使用了很多先进的技术，包括分词、索引和排序算法。",
		},
		{
			"title":   "数据分析平台",
			"content": "大数据时代，数据分析平台变得越来越重要。",
		},
	}

	for i, doc := range documents {
		docJSON, err := json.Marshal(doc)
		if err != nil {
			return err
		}

		req := opensearchapi.IndexRequest{
			Index:      indexName,
			DocumentID: fmt.Sprintf("%d", i+1),
			Body:       bytes.NewReader(docJSON),
		}

		res, err := req.Do(context.Background(), client)
		if err != nil {
			return err
		}
		res.Body.Close()

		if res.IsError() {
			return fmt.Errorf("error indexing document %d: %s", i+1, res.String())
		}
	}

	fmt.Println("Test data inserted successfully")
	return nil
}

func searchDocuments(client *opensearch.Client, queryString string) error {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query":            queryString,
				"fields":           []string{"title", "content"},
				"default_operator": "AND",
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return err
	}

	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(indexName),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return err
		}
		return fmt.Errorf("[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		)
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return err
	}

	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	fmt.Printf("Found %d results:\n", len(hits))
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		fmt.Printf("Title: %s\nContent: %s\n\n", source["title"], source["content"])
	}

	return nil
}

/*
对于包含中文字符、数字等复杂内容的查询，query_string 确实可能是一个更好的选择。query_string 查询提供了更强大和灵活的查询能力，特别适合处理复杂的查询需求。以下是 query_string 的一些优势和特点：

复杂查询语法：支持 AND、OR、NOT 等布尔操作符，允许构建复杂的查询逻辑。

字段指定：可以在查询字符串中指定要搜索的字段。

通配符：支持 * 和 ? 通配符，增加查询的灵活性。

模糊搜索：通过 ~ 操作符支持模糊匹配。

短语搜索：使用引号可以进行精确短语匹配。

范围查询：支持数值和日期范围查询。

权重调整：可以为不同的查询项分配不同的权重。

多字段查询：可以同时在多个字段中搜索。

下面是一个使用 Golang 和 query_string 进行查询的示例：

在这个例子中：

我们使用 query_string 查询，搜索同时包含 "OpenSearch" 和 ("中文" 或 "数字") 的文档。

查询在 "title" 和 "content" 两个字段中进行。

default_operator 设置为 "AND"，意味着查询字符串中的词默认是 AND 关系。

使用 query_string 的优点是它能够处理复杂的查询逻辑，特别适合需要在多个字段中进行复杂搜索的场景。但也需要注意，query_string 查询的语法相对复杂，如果直接暴露给最终用户，可能会带来安全风险，因此通常建议在后端控制查询的构建。

对于中文查询，确保你的索引使用了适当的中文分词器（如 IK 分词器），这样可以提高中文搜索的准确性。
*/

/*
`"tokenizer": "ik_max_word"` 是指定使用 IK 分词器的最大分词模式。IK 分词器是一个流行的中文分词插件，它提供了两种主要的分词模式：

1. `ik_max_word`：最大化分词
2. `ik_smart`：智能分词

让我详细解释一下 `ik_max_word`：

1. 最大化分词模式：
   - 这种模式会将文本切分成尽可能多的词。
   - 它会穷尽各种可能的词语组合。
   - 适合于需要更精细搜索的场景。

2. 工作原理：
   - 它会将一句话分解成所有可能的词语组合。
   - 包括单字、双字词、多字词等。

3. 示例：
   假设我们有一个短语 "中华人民共和国"

   使用 `ik_max_word` 可能会得到如下分词结果：
   - 中华人民共和国
   - 中华人民
   - 中华
   - 华人
   - 人民
   - 共和国
   - 中
   - 华
   - 人
   - 民
   - 共
   - 和
   - 国

4. 优点：
   - 提高召回率：由于生成了所有可能的词组，可以匹配更多的查询。
   - 适合精确搜索：可以找到包含特定词组的文档。

5. 缺点：
   - 可能会产生很多冗余的词，增加索引大小。
   - 在某些情况下可能会影响搜索的精确度。

6. 使用场景：
   - 当你需要非常详细和广泛的搜索结果时。
   - 适合于全文搜索，特别是当用户可能使用文档中的任何部分词语进行搜索时。

相比之下，`ik_smart` 模式会进行更智能的分词，通常只输出较长的词语，不会将词语进一步拆分。它更适合于一般的搜索场景，可以提供更精确但可能召回率较低的结果。

在你的索引设置中使用 `"tokenizer": "ik_max_word"`，意味着你希望在索引和搜索时对文本进行最大化分词，以获得最全面的搜索能力。这对于需要高召回率的中文搜索场景非常有用。

*/

/*
`match_phrase` 和 `query_string` 是 OpenSearch（和 Elasticsearch）中两种不同的查询类型，它们在用途和行为上有显著的区别。让我们详细比较一下：

1. match_phrase

   - 用途：精确短语匹配
   - 特点：
     - 要求查询词组按照指定的顺序完整出现在文档中
     - 不支持复杂的查询语法
     - 通常用于查找包含特定词组的文档

   - 示例：
     ```json
     {
       "query": {
         "match_phrase": {
           "content": "OpenSearch 查询"
         }
       }
     }
     ```
   - 这将匹配包含 "OpenSearch 查询" 这个完整短语的文档

2. query_string

   - 用途：复杂的全文查询
   - 特点：
     - 支持复杂的查询语法，包括 AND、OR、NOT 操作符
     - 可以在多个字段上进行查询
     - 支持通配符、正则表达式、模糊查询等高级功能
     - 允许用户输入更自由的查询表达式

   - 示例：
     ```json
     {
       "query": {
         "query_string": {
           "default_field": "content",
           "query": "OpenSearch AND (查询 OR 搜索)"
         }
       }
     }
     ```
   - 这将匹配包含 "OpenSearch" 且包含 "查询" 或 "搜索" 的文档

主要区别：

1. 精确度 vs 灵活性
   - `match_phrase` 更精确，要求短语完全匹配
   - `query_string` 更灵活，允许复杂的查询逻辑

2. 查询语法
   - `match_phrase` 不支持特殊语法
   - `query_string` 支持丰富的查询语法

3. 使用场景
   - `match_phrase` 适用于需要精确匹配短语的场景
   - `query_string` 适用于需要复杂查询逻辑或允许用户输入高级查询的场景

4. 性能
   - `match_phrase` 通常性能较好，因为它的查询逻辑相对简单
   - `query_string` 可能会较慢，特别是在复杂查询时

5. 安全性
   - `match_phrase` 相对安全，不容易受到注入攻击
   - `query_string` 需要小心处理用户输入，因为它支持复杂语法

6. 字段处理
   - `match_phrase` 通常用于单个字段
   - `query_string` 可以轻松跨多个字段查询

在 Go 中使用这两种查询类型的示例：

```go
// match_phrase 查询
matchPhraseQuery := map[string]interface{}{
    "query": map[string]interface{}{
        "match_phrase": map[string]interface{}{
            "content": "OpenSearch 查询",
        },
    },
}

// query_string 查询
queryStringQuery := map[string]interface{}{
    "query": map[string]interface{}{
        "query_string": map[string]interface{}{
            "default_field": "content",
            "query": "OpenSearch AND (查询 OR 搜索)",
        },
    },
}
```

选择使用哪种查询类型取决于你的具体需求。如果你需要精确的短语匹配，使用 `match_phrase`；如果你需要支持复杂的查询逻辑或用户输入的高级查询，使用 `query_string`。
*/
