bulk的格式：
{action:{metadata}}\n
{requstbody}\n (请求体)

action：(行为)，包含create（文档不存在时创建）、update（更新文档）、index（创建新文档或替换已用文档）、delete（删除一个文档）。
create和index的区别：如果数据存在，使用create操作失败，会提示文档已存在，使用index则可以成功执行。
metadata：(行为操作的具体索引信息)，需要指明数据的_index、_type、_id。
示例：
```
{"delete":{"_index":"lib","_type":"user","_id":"1"}}

{"index":{"_id":1}}  \\行为：索引信息
{"title":"Java","price","55"} \\请求体
```

curl -X POST "localhost:9200/_bulk?pretty" -H 'Content-Type: application/json' -d'
{ "index" : { "_index" : "test", "_id" : "1" } }
{ "field1" : "value1" }
{ "delete" : { "_index" : "test", "_id" : "2" } }
{ "create" : { "_index" : "test", "_id" : "3" } }
{ "field1" : "value3" }
{ "update" : {"_id" : "1", "_index" : "test"} }
{ "doc" : {"field2" : "value2"} }
'



```
curl -X POST "localhost:9200/_bulk?pretty" -H 'Content-Type: application/json' -d'
{"index":{"_index" : "test", "_id":1}}
{"name":"zhang","age": "55"}
{"index":{"_index" : "test", "_id":2}}
{"name":"wang","age": "45"}
{"index":{"_index" : "test", "_id":3}}
{"name":"li","age": "35"}
{"index":{"_index" : "test", "_id":4}}
{"title":"zhao","age": 50}
'
## 直接多个document是错误的
curl -X POST "localhost:9200/_bulk?pretty" -H 'Content-Type: application/json' -d'
{"index":{"_index" : "test"}}
{"name":"zhang11","age": "55"}
{"name":"wang11","age": "45"}
'



curl -X GET "localhost:9200/test/_search?pretty" 得到所有数据，用— _id

```
``` 
//返回结果
{
  "took": 60,
  "error": false //请求是否出错，返回false、具体的错误
  "items": [
     //操作过的文档的具体信息
     {
        "index":{
           "_index": "lib",
           "_type": "user",
           "_id": "1",
           "_version": 1,
           "result": "created", //返回请求结果
           "_shards": {
              "total": 1,
              "successful": 1,
              "failed": 0
           },
           "_seq_no": 0,
           "_primary_trem": 1
           "status": 200
        }
    },
    ...
  ]
}
```

bluk一次最大处理多少数据量
bulk会将要处理的数据载入内存中，所以数据量是有限的，最佳的数据两不是一个确定的数据，它取决于你的硬件，你的文档大小以及复杂性，你的索引以及搜索的负载。
一般建议是1000-5000个文档，大小建议是5-15MB，默认不能超过100M，可以在es的配置文件（即$ES_HOME下的config下的elasticsearch.yml）中，bulk的线程池配置是内核数+1。



```
curl -X POST "localhost:9200/geo_location_sandbox4/_bulk?pretty" -H 'Content-Type: application/json' -d'
{ "index" : { "_id" : "2-2" } }
{"id":2,"kind":2,"snowflakeID":113220548128474149,"name":"外卖餐厅1","removed":false,"cityID":1,"areaID":0,"location":"0.000000,0.000000","address":"","cuisineType":"","cafeteriaID":-1}
'
```