```
curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap?pretty" -H 'Content-Type: application/json' -d'
{
  "aliases": {
    "testaliais": {}
  },
  "mappings": {
    "properties": {
      "id": {"type": "long"},
      "name": { "type": "text", "analyzer": "ik_smart" },
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

curl -X GET "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap/_mapping?pretty"

curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap/_doc/1?pretty" \
-H 'Content-Type: application/json' -d \
'{  "name": "湘潭市中心医院营养食堂（华银）", 
    "id": 1, 
    "remark": "美餐小食堂", 
    "age": 1, 
    "sort": 1, 
    "name2": {"first": "abc-def-aee", "a": 10},
    "tags": ["7ec0e0e5-a4b0-46d7-af56-5b3eab477aea", "aaaabc", "abc"]
}'

curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap/_doc/2?pretty" \
-H 'Content-Type: application/json' -d \
'{  "name": "THE RESTful（武汉顺丰丰泰店）", 
    "id": 1, 
    "remark": "面包天使新乡店", 
    "age": 1, 
    "sort": 1, 
    "name2": {"first": "abc-def-aee", "a": 23},
    "tags": ["7ec0e0e5-a4b0-46d7-af56-5b3eab477ae", "cde", "def"]
}'


curl "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap/_search?pretty"

curl "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "match": {"age": 1} },
   "sort": [
     { "name2.a": "asc" }
   ]
 }
'

// 因为keyword是不分词的，所以是找不到数据的。
curl "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "match": {"name": "7ec0e0e5"} },
   "sort": [
     { "name2.a": "asc" }
   ]
 }
'

curl "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "match": {"remark": "7ec0e0e5"} },
   "sort": [
     { "name2.a": "asc" }
   ]
 }
'


curl "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "term": {"remark.keyword": "7ec0e0e5-a4b0-46d7-af56-5b3eab477aea"} },
   "sort": [
     { "name2.a": "asc" }
   ]
 }
'




curl -X DELETE "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap?pretty"


curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/test?pretty" -H 'Content-Type: application/json' -d'
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
curl "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "terms": { "tags": ["abc", "7ec0e0e5-a4b0-46d7-af56-5b3eab477aea"] }},
   "sort": [
     { "name2.a": "asc" }
   ]
 }
'

## 这样，两个都会查出来，所以是 或查询
curl "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
 {
   "query": { "terms": { "tags": ["abc", "7ec0e0e5-a4b0-46d7-af56-5b3eab477ae"] }},
   "sort": [
     { "name2.a": "asc" }
   ]
 }
'

curl "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
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


curl "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap/_search?pretty" -H 'Content-Type: application/json' -d '
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
curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap?pretty" -H 'Content-Type: application/json' -d'
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


## 中文分词

"analyzer": "ik_smart"

curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/testmap?pretty" -H 'Content-Type: application/json' -d'
{
  "aliases": {
    "testaliais": {}
  },
  "mappings": {
    "properties": {
      "id": {"type": "long"},
      "name": { "type": "text", "analyzer": "ik_smart" },
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