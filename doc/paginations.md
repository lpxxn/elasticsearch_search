https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-body.html#request-body-search-scroll
# from size
```
curl -X GET "localhost:9200/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "from" : 0, "size" : 10,
    "query" : {
        "term" : { "user" : "kimchy" }
    }
}
'
```
>Note that from + size can not be more than the index.max_result_window index setting, which defaults to **10,000**.
可以max_result_window 的值调高 50000。但不好，看下面。
# scroll

While a search request returns a single “page” of results, the scroll API can be used to retrieve large numbers of results (or even all results) from a single search request, in much the same way as you would use a cursor on a traditional database.

>**Scrolling is not intended for real time user requests,** but rather for processing large amounts of data, e.g. in order to reindex the contents of one index into a new index with a different configuration.
```
curl -X POST "localhost:9200/twitter/_search?scroll=1m&pretty" -H 'Content-Type: application/json' -d'
{
    "size": 100,
    "query": {
        "match" : {
            "title" : "elasticsearch"
        }
    }
}
'
```

```
curl -X POST "localhost:9200/_search/scroll?pretty" -H 'Content-Type: application/json' -d'
{
    "scroll" : "1m", 
    "scroll_id" : "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAD4WYm9laVYtZndUQlNsdDcwakFMNjU1QQ==" 
}
'

```
The result from the above request includes a _scroll_id, which should be passed to the scroll API in order to retrieve the next batch of results.

## Sliced Scrolledit
   For scroll queries that return a lot of documents it is possible to split the scroll in multiple slices which can be consumed independently:
```
GET /twitter/_search?scroll=1m
{
    "slice": {
        "id": 0, 
        "max": 2 
    },
    "query": {
        "match" : {
            "title" : "elasticsearch"
        }
    }
}
```   
# Search Afteredit
scan 不推荐用于时间数据，因为他是把数据放到内存中，如果这时对数据进行修改，不会实时反映。
search_afeter 可以对应实时
大于5.0版本，用 search_after

Pagination of results can be done by using the from and size but the cost becomes prohibitive when the deep pagination is reached. The index.max_result_window which defaults to 10,000 is a safeguard, search requests take heap memory and time proportional to from + size. The Scroll api is recommended for efficient deep scrolling but scroll contexts are costly and it is not recommended to use it for real time user requests. The search_after parameter circumvents this problem by providing a live cursor. The idea is to use the results from the previous page to help the retrieval of the next page.

```
curl -X GET "localhost:9200/twitter/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "size": 10,
    "query": {
        "match" : {
            "title" : "elasticsearch"
        }
    },
    "sort": [
        {"date": "asc"},
        {"tie_breaker_id": "asc"}      
    ]
}
'

curl -X GET "localhost:9200/twitter/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "size": 10,
    "query": {
        "match" : {
            "title" : "elasticsearch"
        }
    },
    "search_after": [1463538857, "654323"],
    "sort": [
        {"date": "asc"},
        {"tie_breaker_id": "asc"}
    ]
}
'


```
https://blog.csdn.net/u011228889/article/details/79760167#searchafter-%E7%9A%84%E6%96%B9%E5%BC%8F

