https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-stats.html
## create index
```
curl -XPUT "http://localhost:9200/tmapping1"
```
or 
```
curl -X PUT "localhost:9200/tmapping1?pretty" -H 'Content-Type: application/json' -d'
{
    "settings" : {
        "number_of_shards" : 1
    },
    "mappings" : {
        "properties" : {
            "field1" : { "type" : "text" }
        }
    }
}
'

```
分片数number_of_shards和副本数number_of_replicas

Changes an index setting in real time.
```
curl -X PUT "localhost:9200/tmapping1/_settings?pretty" -H 'Content-Type: application/json' -d'
{
    "index" : {
        "number_of_replicas" : 2
    }
}
'
```
Adds new fields to an existing index or changes the search settings of existing fields.

```
curl -X PUT "http://localhost:9200/tmapping1/_mapping/?pretty" -H 'Content-Type: application/json' -d '
{
    "properties": {
        "pid":{
        "type": "keyword"
        }
    }
}
'

```

### get mapping
```
curl -X GET "localhost:9200/tmapping1/_mapping?pretty"

curl -X GET "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/my_test_3/_mapping?pretty"
```


## demo

### Use the create index API to create an index with the city text field.

```
curl -X PUT "localhost:9200/my_index?pretty" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "properties": {
      "city": {
        "type": "text"
      }
    }
  }
}
'


```

or if the index is exist 
```
curl -X PUT "http://localhost:9200/my_index/_mapping/?pretty" -H 'Content-Type: application/json' -d '
{
    "properties": {
        "city":{
        "type": "text"
        }
    }
}
'

curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/my_test_3/_mapping?pretty" -H 'Content-Type: application/json' -d'
{
  "properties": {
    "city": {
      "type": "text"
    },
    "location": {
      "type": "geo_point"
    }
  }
}
'

```

## get mapping
```
curl -X GET "localhost:9200/my_index/_mapping?pretty"

{
  "my_index" : {
    "mappings" : {
      "properties" : {
        "city" : {
          "type" : "text"
        },
        "name" : {
          "properties" : {
            "first" : {
              "type" : "text"
            },
            "last" : {
              "type" : "text"
            }
          }
        }
      }
    }
  }
}
```

While text fields work well for full-text search, keyword fields are not analyzed and may work better for sorting or aggregations.

Use the put mapping API to enable a multi-field for the city field. This request adds the city.raw keyword multi-field, which can be used for sorting.

```
curl -X PUT "localhost:9200/my_index/_mapping?pretty" -H 'Content-Type: application/json' -d'
{
  "properties": {
    "city": {
      "type": "text",
      "fields": {
        "raw": {
          "type": "keyword"
        }
      }
    }
  }
}
'

```

### get mapping

```
curl -X GET "localhost:9200/my_index/_mapping?pretty"

{
  "my_index" : {
    "mappings" : {
      "properties" : {
        "city" : {
          "type" : "text",
          "fields" : {
            "raw" : {
              "type" : "keyword"
            }
          }
        },
```

### Add new properties to an existing object fieldedit
You can use the put mapping API to add new properties to an existing object field. To see how this works, try the following example.

Use the create index API to create an index with the name object field and an inner first text field.
```
curl -X PUT "localhost:9200/my_index?pretty" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "properties": {
      "name": {
        "properties": {
          "first": {
            "type": "text"
          }
        }
      }
    }
  }
}
'

```

Use the put mapping API to add a new inner last text field to the name field.

```
curl -X PUT "localhost:9200/my_index/_mapping?pretty" -H 'Content-Type: application/json' -d'
{
  "properties": {
    "name": {
      "properties": {
        "last": {
          "type": "text"
        }
      }
    }
  }
}
'

```

Use the get mapping API to verify your changes.
```
curl -X GET "localhost:9200/my_index/_mapping?pretty"

```
