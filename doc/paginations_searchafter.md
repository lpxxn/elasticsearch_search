

https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn
```
## get mapping
curl -X GET "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/geo_location_sandbox4/_mapping?pretty"
## query
curl -X GET "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/geo_location_sandbox4/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "query": {
        "term" : {
            "snowflakeID" : 51037171290243083
        }
    }
}
'
curl -X GET "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/geo_location_sandbox4/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "size": 10,
    "query": {
        "match_phrase" : {
            "name" : "食堂"
        }
    },
    "sort": [
        {"snowflakeID": "desc"}
    ]
}
'
## search after
curl -X GET "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/geo_location_sandbox4/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "size": 10,
    "query": {
        "match_phrase" : {
            "name" : "食堂"
        }
    },
    "sort": [
        {"snowflakeID": "asc"}
    ]
}
'

curl -X GET "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/geo_location_sandbox4/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "size": 10,
    "query": {
      "bool": {
        "filter": {
            "match": {
              "kind": 1
            }
        },
        "must": {
            "match_phrase": {
              "name": "食堂"
            }
        }
      }
    },
    "sort": [
        {"snowflakeID": "desc"}
    ]
}
'

curl -X GET "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/geo_location_sandbox4/_search?pretty" -H 'Content-Type: application/json' -d'
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
curl --location --request GET 'https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/geo_location_sandbox4/_search' \
--header 'Content-Type: application/json' \
--data '{"query": { "term": { "kind": 2}},"size": 0,"track_total_hits": true}'

curl --location --request GET 'https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/geo_location_sandbox4/_search' \
--header 'Content-Type: application/json' \
--data '{"query": { "term": { "kind": 1}},"size": 0,"track_total_hits": true}'

curl --location --request GET 'https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/geo_location_sandbox4/_count'