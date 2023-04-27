

https://hello.es.amazonaws.com.cn
```
## get mapping
curl -X GET "https://hello.es.amazonaws.com.cn/geo_location_sandbox4/_mapping?pretty"
## query
curl -X GET "https://hello.es.amazonaws.com.cn/geo_location_sandbox4/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "query": {
        "term" : {
            "snowflakeID" : 51037171290243083
        }
    }
}
'

## search after
curl -X GET "https://hello.es.amazonaws.com.cn/geo_location_sandbox4/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "size": 10,
    "query": {
        "term" : {
            "name" : "小"
        }
    },
    "sort": [
        {"snowflakeID": "asc"}
    ]
}
'

curl -X GET "https://hello.es.amazonaws.com.cn/geo_location_sandbox4/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "size": 10,
    "query": {
      "bool": {
        "filter": {
            "match": {
              "kind": 2
            }
        },
        "must": {
            "term": {
              "name": "小"
            }
        }
      }
    },
    "sort": [
        {"snowflakeID": "asc"}
    ]
}
'

curl -X GET "https://hello.es.amazonaws.com.cn/geo_location_sandbox4/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "size": 10,
    "query": {
      "bool": {
        "filter": {
            "match": {
              "kind": 2
            }
        },
        "must": {
            "term": {
              "name": "小"
            }
        }
      }
    },
    "search_after": [62971038981461058],
    "sort": [
        {"snowflakeID": "asc"}
    ]
}
'

```