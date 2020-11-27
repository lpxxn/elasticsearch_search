```
curl -X PUT "localhost:9200/testmap?pretty" -H 'Content-Type: application/json' -d'
{
  "aliases": {
    "testaliais": {}
  },
  "mappings": {
    "properties": {
      "id": {"type": "long"},
      "name": { "type": "keyword" },
      "remart": {"type": "text"},
      "age": {"type": "integer"},
      "sort": {"type": "integer"},
      "name2": {
        "properties": {
          "first": {
            "type": "text"
          },
          "a": {
              "type": "integer"
          }
        }
      },
      "tags": {"type": "keyword"}
    }
  }
}
'

类型是keyword,就是不分词处理

curl -X GET "localhost:9200/testmap/_mapping?pretty"

curl -X PUT "localhost:9200/testmap/_doc/1?pretty" \
-H 'Content-Type: application/json' -d \
'{  "name": "7ec0e0e5-a4b0-46d7-af56-5b3eab477aea", 
    "id": 1, 
    "remark": "7ec0e0e5-a4b0-46d7-af56-5b3eab477aea", 
    "age": 1, 
    "sort": 1, 
    "name2": {"first": "abc-def-aee", "a": 10},
    "tags": ["7ec0e0e5-a4b0-46d7-af56-5b3eab477aea", "aaaabc", "abc"]
}'

curl -X PUT "localhost:9200/testmap/_doc/2?pretty" \
-H 'Content-Type: application/json' -d \
'{  "name": "7ec0e0e5-a4b0-46d7-af56-5b3eab477ae", 
    "id": 1, 
    "remark": "7ec0e0e5-a4b0-46d7-af56-5b3eab477ae", 
    "age": 1, 
    "sort": 1, 
    "name2": {"first": "abc-def-aee", "a": 23},
    "tags": ["7ec0e0e5-a4b0-46d7-af56-5b3eab477ae", "cde", "def"]
}'


curl "localhost:9200/testmap/_search?pretty"

curl "localhost:9200/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "match": {"age": 1} },
   "sort": [
     { "name2.a": "asc" }
   ]
 }
'

// 因为keyword是不分词的，所以是找不到数据的。
curl "localhost:9200/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "match": {"name": "7ec0e0e5"} },
   "sort": [
     { "name2.a": "asc" }
   ]
 }
'

curl "localhost:9200/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "match": {"remark": "7ec0e0e5"} },
   "sort": [
     { "name2.a": "asc" }
   ]
 }
'


curl "localhost:9200/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "term": {"remark.keyword": "7ec0e0e5-a4b0-46d7-af56-5b3eab477aea"} },
   "sort": [
     { "name2.a": "asc" }
   ]
 }
'




curl -X DELETE "localhost:9200/testmap?pretty"


curl -X PUT "localhost:9200/test?pretty" -H 'Content-Type: application/json' -d'
{
  "aliases": {
    "alias_1": {},
    "alias_2": {
      "filter": {
        "term": { "user.id": "kimchy" }
      },
      "routing": "shard-1"
    }
  }
}
'


```
##  array 查询

```
## array 还是keyword比较好
curl "localhost:9200/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "terms": { "tags": ["abc", "7ec0e0e5-a4b0-46d7-af56-5b3eab477aea"] }},
   "sort": [
     { "name2.a": "asc" }
   ]
 }
'

## 这样，两个都会查出来，所以是 或查询
curl "localhost:9200/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "terms": { "tags": ["abc", "7ec0e0e5-a4b0-46d7-af56-5b3eab477ae"] }},
   "sort": [
     { "name2.a": "asc" }
   ]
 }
'

curl "localhost:9200/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { 
        "bool": {
            "should": [
                {"terms": { "tags": ["abc", "7ec0e0e5-a4b0-46d7-af56-5b3eab477ae"]}},
                {"term": {"age": 1}}
            ]
        }
    },
   "sort": [
     { "name2.a": "asc" }
   ]
 }
'


curl "localhost:9200/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": {
       "term": {"tags": "7ec0e0e5-a4b0-46d7-af56-5b3eab477ae"}
    },
   "sort": [
     { "name2.a": "asc" }
   ]
 }
'

```

## error
```
curl -X PUT "localhost:9200/testmap?pretty" -H 'Content-Type: application/json' -d'
{
  "aliases": {
    "testaliais": {}
  },
  "mappings": {
    "properties": {
      "id": {"type": "long"},
      "name": { "type": "keyword" },
      "remart": {"type": "text"},
      "age": {"type": "integer"},
      "sort": {"type": "integer"},
      "name2": {
          // 不加 properties 会报错
          "first": {
            "type": "text"
          },
          "a": {
              "type": "integer"
          }
      }
    }
  }
}
'

```
