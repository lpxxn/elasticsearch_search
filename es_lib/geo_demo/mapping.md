https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-geo-distance-query.html
https://www.elastic.co/guide/cn/elasticsearch/guide/current/lat-lon-formats.html

```
curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1?pretty" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "properties": {
      "hotel": {
        "properties": {
          "name" : { "type" : "text" },
          "location": {
            "type": "geo_point"
          }
        }
      }
    }
  }
}
'

curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1/_doc/1?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "Chipotle Mexican Grill",
      "location": "40.715, -74.011" 
    }
}
'


curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1/_doc/2?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
    "name": "速8",
    "location": {
      "lat": 40.12,
      "lon": -71.34
    }
  }
}
'

curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1/_doc/3?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "Pala Pizza",
      "location": { 
        "lat":     40.722,
        "lon":    -73.989
      }
    }
}
'

curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1/_doc/4?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "Mini Munchies Pizza",
      "location": [ -73.983, 40.719 ] 
    }
}
'

```

## 通过地理坐标点过滤
这是目前为止最有效的地理坐标过滤器了，因为它计算起来非常简单。 你指定一个矩形的 顶部 , 底部 , 左边界 ，和 右边界 ，然后过滤器只需判断坐标的经度是否在左右边界之间，纬度是否在上下边界之间：
```
curl -X GET "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": {
    "bool": {
      "must": {
        "match_all": {}
      },
      "filter": {
        "geo_bounding_box": {
          "hotel.location": {
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
distance

```
curl -X GET "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": {
    "bool": {
      "must": {
        "match_all": {}
      },
      "filter": {
        "geo_distance": {
          "distance": "200km",
          "hotel.location": {
            "lat": 40,
            "lon": -70
          }
        }
      }
    }
  }
}
'
```

```
curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/my_geoshapes?pretty" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "properties": {
      "hotel": {
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
curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/my_geoshapes/_doc/1?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
    "location": {
      "type" : "polygon",
      "coordinates" : [[[13.0 ,51.5], [15.0, 51.5], [15.0, 54.0], [13.0, 54.0], [13.0 ,51.5]]]
    }
  }
}
'

curl "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/my_geoshapes/_search?pretty"

curl -X GET "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/my_geoshapes/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": {
    "bool": {
      "must": {
        "match_all": {}
      },
      "filter": {
        "geo_bounding_box": {
          "hotel.location": {
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

To match both geo_point and geo_shape values, search both indices:
```
curl -X GET "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1,my_geoshapes/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": {
    "bool": {
      "must": {
        "match_all": {}
      },
      "filter": {
        "geo_distance": {
          "distance": "200km",
          "hotel.location": {
            "lat": 40,
            "lon": -70
          }
        }
      }
    }
  }
}
'

curl -X GET "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1,my_geoshapes/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": {
    "bool": {
      "must": {
        "match_all": {}
      },
      "filter": {
        "geo_distance": {
          "distance": "200km",
          "hotel.location": [ -70, 40 ]
        }
      }
    }
  }
}
'

curl -X GET "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1,my_geoshapes/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": {
    "bool": {
      "must": {
        "match_all": {}
      },
      "filter": {
        "geo_bounding_box": {
          "hotel.location": {
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
不添加新功能的前提下，重构代码，调整原有的程序逻辑，以适应新功能实现的功能。
再