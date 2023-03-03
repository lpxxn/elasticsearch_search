https://www.elastic.co/guide/en/elasticsearch/reference/7.10/query-dsl-geo-bounding-box-query.html
7.10是不支持 geo_shape 的geo_bounding_box 查询
7.11 版本之后才能用了。
```

curl -X PUT "localhost:9200/my_locations2?pretty" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "properties": {
      "pin": {
        "properties": {
          "location": {
            "type": "geo_point"
          }
        }
      }
    }
  }
}
'
curl -X PUT "localhost:9200/my_locations2/_doc/1?pretty" -H 'Content-Type: application/json' -d'
{
  "pin": {
    "location": {
      "lat": 40.12,
      "lon": -71.34
    }
  }
}
'


curl -X PUT "localhost:9200/my_geoshapes2?pretty" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "properties": {
      "pin": {
        "properties": {
          "location": {
            "type": "geo_shape"
          }
        }
      }
    }
  }
}
'
curl -X PUT "localhost:9200/my_geoshapes2/_doc/1?pretty" -H 'Content-Type: application/json' -d'
{
  "pin": {
    "location": {
      "type" : "polygon",
      "coordinates" : [[[13.0 ,51.5], [15.0, 51.5], [15.0, 54.0], [13.0, 54.0], [13.0 ,51.5]]]
    }
  }
}
'

curl "localhost:9200/my_geoshapes2/_search?pretty"

curl -X GET "localhost:9200/my_geoshapes2/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": {
    "bool": {
      "must": {
        "match_all": {}
      },
      "filter": {
        "geo_bounding_box": {
          "pin.location": {
            "top_left": {
              "lat": 40.73,
              "lon": -74.1
            },
            "bottom_right": {
              "lat": 40.01,
              "lon": -71.12
            }
          }
        }
      }
    }
  }
}
'


```