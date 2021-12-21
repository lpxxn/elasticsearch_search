```
curl -X PUT "localhost:9200/mytest_geo1/_doc/1?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "美餐小食堂",
      "location": "39.97624, 116.49174" 
    }
}
'

// 40 米
curl -X PUT "localhost:9200/mytest_geo1/_doc/2?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "40米食堂A",
      "location": "39.97591, 116.49178" 
    }
}
'

// 40 米2
curl -X PUT "localhost:9200/mytest_geo1/_doc/3?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "40米食堂B",
      "location": "39.97658, 116.49179" 
    }
}
'

// 50 米
curl -X PUT "localhost:9200/mytest_geo1/_doc/4?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "50米食堂A",
      "location": "39.97625, 116.49118" 
    }
}
'

// 100 米
curl -X PUT "localhost:9200/mytest_geo1/_doc/5?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "100米食堂A",
      "location": "39.97643, 116.49045" 
    }
}
'

// 外
curl -X PUT "localhost:9200/mytest_geo1/_doc/6?pretty" -H 'Content-Type: application/json' -d'
{
  "hotel": {
      "name":     "大山子798A",
      "location": "39.97778, 116.48992" 
    }
}
'

curl -X PUT "localhost:9200/mytest_geo1/_doc/7?pretty" -H 'Content-Type: application/json' -d'
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
curl -X GET "localhost:9200/mytest_geo1/_search?pretty" -H 'Content-Type: application/json' -d'
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
在做开发时，合理的扩展是正常的，为了防止未来变更过度设计，我们，因为我们并不确定未来的需求变动。