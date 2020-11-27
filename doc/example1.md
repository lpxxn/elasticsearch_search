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
      }
    }
  }
}
'

类型是keyword,就是不分词处理

curl -X GET "localhost:9200/testmap/_mapping?pretty"

curl -X PUT "localhost:9200/testmap/_doc/1?pretty" \
-H 'Content-Type: application/json' -d'{"name": "7ec0e0e5-a4b0-46d7-af56-5b3eab477aea", "id": 1, "remark": "7ec0e0e5-a4b0-46d7-af56-5b3eab477aea", "age": 1, "sort": 1, "name2": {"first": "abc-def-aee", "a": 10}}'

curl "localhost:9200/testmap/_search?pretty"

curl "localhost:9200/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "match": {"age": 1} },
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
