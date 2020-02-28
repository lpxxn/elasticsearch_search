
#### delete index
```
curl -X DELETE "http://localhost:9200/some_index?pretty"
```

all data:
curl "http://localhost:9200/my_test_3/_mapping?pretty"


#### delete index document
DELETE /index/doc/id
```
curl -XDELETE "http://localhost:9200/mytest2/_doc/iQSmdXABXv3MPkS7Cg-Z?pretty"

```

#####delete all
```
curl -XPOST "http://localhost:9200/my_test_3/_delete_by_query?conflicts=proceed&pretty" -H 'Content-Type: application/json' -d '
{
  "query": {
    "match_all": {}
  }
}
'
```

#### delete by query
```
curl -XPOST "http://localhost:9200/my_test_3/_delete_by_query?conflicts=proceed&pretty" -H 'Content-Type: application/json' -d '
{
  "query": {
    "term": {"age": 1}
  }
}
'
```

```
curl -XPOST "http://localhost:9200/my_test_3/_delete_by_query?conflicts=proceed&pretty" -H 'Content-Type: application/json' -d '
{
  "query": {
    "term": {"age" : 1 }
  }
}
'
```


