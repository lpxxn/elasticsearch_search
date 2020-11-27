
## Filter
过滤器(Filters)和缓存 有时间查一下
在执行filter和query时,先执行filter在执行query
查询(query)和过滤(filter)。查询即是之前提到的query查询，它 (查询)默认会计算每个返回文档的得分，然后根据得分排序。而过滤(filter)只会筛选出符合的文档，并不计算 得分，且它可以缓存文档 。所以，单从性能考虑，过滤比查询更快。
原则上来说，使用查询语句做全文本搜索或其他需要进行相关性评分的时候，剩下的全部用过滤语句

换句话说，过滤适合在大范围筛选数据，而查询则适合精确匹配数据。一般应用时， 应先使用过滤操作过滤数据， 然后使用查询匹配数据
Elasticsearch会自动缓存经常使用的过滤器，以加快性能。
https://www.elastic.co/guide/en/elasticsearch/reference/current/query-filter-context.html
Frequently used filters will be cached automatically by Elasticsearch, to speed up performance.


```
curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
    "query" : {
        "bool" : {
            "filter" : {
                "range" : {
                    "age" : { "gt" : 1 }
                }
            },
            "must" : {
                "match" : {
                    "remark" : "三"
                }
            }
        }
    }
}
'

curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
    "query" : {
        "bool" : {
            "filter" : {
                "range" : {
                    "age" : { "gt" : 1 }
                }
            }
        }
    }
}
'

curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
    "query" : {
        "bool" : {
            "filter" : {
                "match" : {
                    "age" : 1
                }
            }
        }
    }
}
'
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-constant-score-query.html
If we want to speed up term query and get it cached then it should be wrapped up in a constant_score filter.
curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
   "query": {
	    "constant_score": {
	        "filter" : {
				"match": {"age": 18}
		}
   }
}
'
```

```
curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
    "query": {
        "bool": {
            "filter": {
                "match": {"remark" : "三"}
            },
            "should": [
                {"match": {"age" : 1}},
                {"match": {"age" : 5}}
            ]
            
        }
    }
}
'

curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
    "query": {
        "bool": {
            "filter": [
                {"match": {"remark" : "三"}},
                {"match": {"age" : 1}}
            ]
        }
    }
}
'
```

terms 过滤
```
{
  "query": {
    "bool": {
      "must": [
        {"term": {
          "name": {
            "value": "中国"
          }
        }}
      ],
      "filter": {
        "terms": {
          "content":[
              "科技",
              "声音"
            ]
        }
      }
    }
  }
}
```
