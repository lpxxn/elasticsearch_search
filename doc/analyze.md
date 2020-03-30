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

但是在es5.0以上版本的可以通过在filed增加keyword就可以查询到，因为text类型数据会创建两份索引，其中一份是长度为256的keyword索引数据

curl "http://localhost:9200/my_test_3/_search?pretty" -H 'Content-Type: application/json' -d '
{
  "query": {
        "term": {
          "pid.keyword": "9z80e0e5-b9b0-46d7-af56-5b3eab47icea"
        }
  }
}
'
```
当然另外一种方式就是创建mapping,指定pID的类型是keyword,就是不分词处理，但是这个需要在我们往index 中插入数据之前，一旦插入了数据，是不能在创建mapping的，只能通过reindex重新数据迁移。
```
curl -XPUT "http://localhost:9200/my_test_3/?pretty" -H 'Content-Type: application/json' -d '
{
"mappings": {
 "users": {
   "properties": {
     "pid":{
       "type": "keyword"
     }
   }
 }
}
}
'
如果已经存在会返回错误
resource_already_exists_exception

详细可以查看 mapping.md