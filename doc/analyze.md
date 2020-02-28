```
curl -XPOST "http://localhost:9200/my_test_3/_analyze?pretty" -H 'Content-Type: application/json' -d '
{
  "text":"12一二"
}
'
```

分析 发现会被生成4部分
```
curl -XPOST "http://localhost:9200/my_test_3/_analyze?pretty" -H 'Content-Type: application/json' -d '
{
  "text":"7ec0e0e5-a4b0-46d7-af56-5b3eab477aea"
}
'
```
