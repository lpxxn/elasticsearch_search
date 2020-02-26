
## api

### info 
```
curl GET -v 'localhost:9200'
```

default username and pwd
```
user: elastic
password: changeme
```
修改密码需要加上 `-e "ELASTIC_PASSWORD=my_own_password"`

```
curl -u elastic:changeme localhost:9200   
```

```                                                                                           
{
  "name" : "c8a3eb9ff9a1",
  "cluster_name" : "docker-cluster",
  "cluster_uuid" : "CQU8GIFFS0Kps3LMYRWxTw",
  "version" : {
    "number" : "7.6.0",
    "build_flavor" : "default",
    "build_type" : "docker",
    "build_hash" : "7f634e9f44834fbc12724506cc1da681b0c3b1e3",
    "build_date" : "2020-02-06T00:09:00.449973Z",
    "build_snapshot" : false,
    "lucene_version" : "8.4.0",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
```

```
curl -X GET "localhost:9200/_cat/nodes?v&pretty"
```

## get all index
```
curl "localhost:9200/_cat/indices?v"

curl "http://localhost:9200/_aliases?pretty=true"

# 在 indices 里
curl "localhost:9200/_stats?pretty=true"


```

### create an index
create an index called mytest 
```
curl -X PUT "localhost:9200/mytest?pretty"
```

```
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "mytest"
}
```

delete
```
curl -X DELETE "localhost:9200/twitter?pretty"

```
#### mapping
```
curl -v -XPUT "http://localhost:9200/some_index?pretty" -H 'Content-Type: application/json' -d ' 
{
  "mappings": {
    "properties": {
      "field_1": { "type" : "text" },
        "field_2": { "type" : "integer" },
        "field_3": { "type" : "boolean" },
      "created": {
        "type" : "date",
        "format" : "epoch_second"
      }
    }
  }
}
'
```
#### delete index
```
curl -X DELETE "http://localhost:9200/some_index?pretty"
```

#### index stats
```
curl -X GET "localhost:9200/some_index/_stats?pretty"

```

#### delete index document
DELETE /index/doc/id
```
curl -XDELETE "http://localhost:9200/mytest2/_doc/iQSmdXABXv3MPkS7Cg-Z?pretty"

curl -XPOST "http://localhost:9200/my_test_3/_delete_by_query?conflicts=proceed&pretty" -H 'Content-Type: application/json' -d '
{
  "query": {
    "match_all": {}
  }
}
'

```

### add a new document

add a document to the mytest index:
```
curl -X PUT "localhost:9200/mytest/_doc/1?pretty" \
-H 'Content-Type: application/json' -d'{"name": "Mark Heath" }'


curl -X PUT "localhost:9200/some_index/people/1?pretty" \
-H 'Content-Type: application/json' -d'{"name": "Mark Heath" }'
```

### view documents in the index

```
curl "localhost:9200/mytest/_search?pretty"

```
### cluster health
```
curl "http://localhost:9200/_cluster/health?pretty"
```

### Search
simple
```
// 查询的是具体的

curl "http://localhost:9200/my_test_3/_search?q=remark:三&pretty"
curl "http://localhost:9200/my_test_3/_search?q=age:1&pretty"
```
```
curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "query": { "match_all": {} },
  "sort": [
    { "age": "asc" }
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

```

#### 分页
每个搜索请求都是独立的：Elasticsearch 不维护任何请求中的状态信息。如果做分页的话，请在请求中指定 From 和 Size 参数。
如图所示：搜索第 10 到 19 的数据

```
curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "query": { "match_all": {} },
  "sort": [
    { "age": "asc" }
  ],
  "from": 0,
  "size": 10
}
'

```

### match
#### 只要 match 
```
curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "match": {"age": 1} },
   "sort": [
     { "age": "asc" }
   ]
 }
 '

curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "query": { "match": {"name": "wang_cai xiaohei"} }
}
'

curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "query": { "match": {"remark": "一 三"} }
}
'

{
  "query": { "match": {"remark": "测"} }
}

curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "query": { "match": {"remark": "一 三"} }
}
'

curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "query": { 
    "match": {
        "remark": {
            "query": "一 三",
            "operator": "and"
        }
    }
  }
}
'

```
#### term

```
curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "query": {
        "term": {
          "pid": "9z80e0e5-b9b0-46d7-af56-5b3eab47icea"
        }
  }
}
'
上面的语句是无法搜索的到的，因为通过使用分词分析，9z80e0e5-b9b0-46d7-af56-5b3eab47icea会被分拆成4个部分建立倒排索引
分析 发现会被生成4部分
curl -XPOST "http://localhost:9200/my_test_3/_analyze?pretty" -H 'Content-Type: application/json' -d '
{
  "text":"7ec0e0e5-a4b0-46d7-af56-5b3eab477aea"
}
'
但是在es5.0以上版本的可以通过在filed增加keyword就可以查询到，因为text类型数据会创建两份索引，其中一份是长度为256的keyword索引数据

curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "query": {
        "term": {
          "pid.keyword": "9z80e0e5-b9b0-46d7-af56-5b3eab47icea"
        }
  }
}
'
```
当然另外一种方式就是创建mapping,指定pID的类型是keyword,就是不分词处理，但是这个需要在我们往index 中插入数据之前，一旦插入了数据，是不能在创建mapping的，只能通过reindex重新数据迁移。
```
curl -XPUT "http://localhost:9200/my_test_3/?pretty" -H 'Content-Type: application/json' -d '
{
"mappings": {
 "users": {
   "properties": {
     "pid":{
       "type": "keyword"
     }
   }
 }
}
}
'
如果已经存在会返回错误
resource_already_exists_exception

curl -XPUT "http://localhost:9200/my_test_3/_mapping/users?pretty" -H 'Content-Type: application/json' -d '
{
   "properties": {
     "pid":{
       "type": "keyword"
     }
   }
}
'
curl "http://localhost:9200/my_test_3/_mapping?pretty"
https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html

```

#### terms
terms 像sql里的 IN

```
curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
    "query" : {
        "terms" : {
            "age" : [1, 2, 5]
        }
    }
}
'
```
### range
like sql between
gt is greater than
gte is greater than or equal to
lt is less than
lte is less than or equal to

```
curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
    "query": {
        "range" : {
            "age" : {
                "gte" : 4,
                "lte" : 17
            }
        }
    }
}
'
```

#### bool
布尔查询中每个 must，should,must_not 都被称为查询子句。每个
must 或者 should 查询子句中的条件都会影响文档的相关得分。得分越高，文档跟搜索条件匹配得越好。默认情况下，Elasticsearch 返回的文档会根据相关性算分倒叙排列。

must_not 子句中认为是过滤条件。它会过滤返回结果，但不会影响文档的相关性算分，你还可以明确指定任意过滤条件去筛选结构化数据文档。

```
curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{ 
  "query": {
     "bool": {
       "must": [
         { "match": { "remark": "三" } }
       ],
       "must_not": [
         { "match": { "age": 1 } }
       ]
     }
   }
 }
'

curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "query": {
    "bool": {
      "must": { "match_all": {} },
      "filter": {
        "range": {
          "age": {
            "gte": 2,
            "lte": 10 
          }
        }
      }
    }
  }
}
'

```

#### filter
过滤器(Filters)和缓存 有时间查一下
在执行filter和query时,先执行filter在执行query
查询(query)和过滤(filter)。查询即是之前提到的query查询，它 (查询)默认会计算每个返回文档的得分，然后根据得分排序。而过滤(filter)只会筛选出符合的文档，并不计算 得分，且它可以缓存文档 。所以，单从性能考虑，过滤比查询更快。

换句话说，过滤适合在大范围筛选数据，而查询则适合精确匹配数据。一般应用时， 应先使用过滤操作过滤数据， 然后使用查询匹配数据
Elasticsearch会自动缓存经常使用的过滤器，以加快性能。
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


#### match_phrase 短语要在一起
不能查询多个
```
curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "query": { "match_phrase": {"remark": "测试"} }
}
'
```

#### exist

```
{
    "query": {
        "exists": {
            "field": "<your_field_name>"
        }
    }
}
```


### aggregation 聚合
```
curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "aggs": {
    "all_ages": {
      "terms": { "field": "age" }
    }
  }
}
'
```

```
curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "query": {
    "match": {
      "remark": "一 二 花"
    }
  },
  "aggs": {
    "all_ages": {
      "terms": { "field": "age" }
    }
  }
}
'
```

## books
https://wiki.jikexueyuan.com/project/elasticsearch-definitive-guide-cn/010_Intro/10_Installing_ES.html
https://learnku.com/docs/elasticsearch73/7.3
https://es.xiaoleilu.com/index.html


