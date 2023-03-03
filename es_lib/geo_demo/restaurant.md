```
curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1/_doc/1?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "美餐小食堂",
      "location": "39.97624, 116.49174" 
    }
}
'

// 40 米
curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1/_doc/2?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "40米食堂A",
      "location": "39.97591, 116.49178" 
    }
}
'

// 40 米2
curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1/_doc/3?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "40米食堂B",
      "location": "39.97658, 116.49179" 
    }
}
'

// 50 米
curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1/_doc/4?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "50米食堂A",
      "location": "39.97625, 116.49118" 
    }
}
'

// 100 米
curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1/_doc/5?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "100米食堂A",
      "location": "39.97643, 116.49045" 
    }
}
'

// 外
curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1/_doc/6?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "大山子798A",
      "location": "39.97778, 116.48992" 
    }
}
'

curl -X PUT "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1/_doc/7?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "大山子798艺术区A",
      "location": "39.980006, 116.48891" 
    }
}
'

box left  39.97665, 116.49036  right: 39.97584, 116.49211

```

box
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
              "lat": 39.97665,
              "lon": 116.49036
            },
            "bottom_right": {
              "lat": 39.97584,
              "lon": 116.49211
            }
          }
        }
      }
    }
  }
}
'
```

// https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-shape.html
Envelopeedit
Elasticsearch supports an envelope type, which consists of coordinates for upper left and lower right points of 
the shape to represent a bounding rectangle in the format [[minLon, maxLat], [maxLon, minLat]]:
```
curl -X GET "https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn/mytest_geo1/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": {
    "bool": {
      "must": {
        "match_all": {}
      },
      "filter": {
        "geo_shape": {
          "hotel.location": {
            "shape": {
              "type": "envelope",
              "coordinates": [ [ 116.49036,39.97665 ], [ 116.49211,39.97584 ] ]
            }
          }
        }
      }
    }
  }
}
'

```
东西经(longitude)，南北纬(latitude)
Polygonedit
A polygon is defined by a list of a list of points. 
The first and last points in each (outer) list must be the same (the polygon must be closed). 
The following is an example of a Polygon in GeoJSON.