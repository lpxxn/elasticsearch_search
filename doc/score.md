```
Elasticsearch 默认是按照文档与查询的相关度(匹配度)的得分倒序返回结果的. 
得分 (_score) 就越大, 表示相关性越高.

curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "query": { "match_all": {}},
  "sort": [
    { "age": "asc" },
    { "name" : "desc" }
  ]
}
'

took – 查询花费时长（毫秒）
timed_out – 请求是否超时
_shards – 搜索了多少文档，成功、失败或者跳过了多个分片（明细）
max_score – 最相关的文档分数
hits.total.value - 找到的文档总数
hits.sort - 文档排序方式 （如没有则按相关性分数排序）
hits._score - 文档的相关性算分 (match_all 没有算分)


// 加上 "track_scores": true, 后 _score 字段会计算，为false时，返回 `null`
curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "track_scores": true,
  "query": { "match_all": {} },
  "sort": [
    { "age": "asc" },
    { "name" : "desc" }
  ]
}
'

curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "track_scores": true,
  "query" : {
        "match" : { "remark" : "一" }
  },
  "sort": [
    { "age": "asc" },
    { "name" : "desc" }
  ]
}
'


// 当有 Fielddata is disabled on text fields by default. Set fielddata=true 错误时
// 已经存在的数据直接修改
// Fielddata can consume a lot of heap space, especially when loading high cardinality text fields.
curl -X PUT "localhost:9200/my_test_3/_mapping?pretty" -H 'Content-Type: application/json' -d'
{
  "properties": {
    "name": { 
      "type":     "text",
      "fielddata": true
    }
  }
}
'

curl "http://localhost:9200/my_test_3/_mapping?pretty"  

```