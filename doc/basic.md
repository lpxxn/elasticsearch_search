在Elasticsearch中存储数据的行为就叫做 `索引(indexing)` ，不过在索引之前，我们需要明确数据应该存储在哪里。
在Elasticsearch中，文档归属于一种`类型(type)`,而这些类型存在于`索引(index)`中，我们可以画一些简单的对比图来类比传统关系型数据库：
```
Relational DB -> Databases -> Tables -> Rows -> Columns
Elasticsearch -> Indices   -> Types  -> Documents -> Fields
```
[从es6开始，一个index 只能有一个type](https://www.elastic.co/guide/en/elasticsearch/reference/6.0/removal-of-types.html)
https://stackoverflow.com/questions/50820309/elasticsearch-6-rejecting-mapping-update-as-the-final-mapping-would-have-more-t    
Elasticsearch集群可以包含多个`索引(indices)（数据库)`        
每一个索引可以包含一个`类型(types)（表）`    
每一个类型包含多个`文档(documents)（行）`     
然后每个文档包含多`个字段(Fields)（列）`  

索引（名词） 如上文所述，一个索引(index)就像是传统关系数据库中的数据库，它是相关文档存储的地方，index的复数是indices 或indexes。

索引（动词） 「索引一个文档」表示把一个文档存储到索引（名词）里，以便它可以被检索或者查询。这很像SQL中的INSERT关键字，差别是，如果文档已经存在，新的文档将覆盖旧的文档。

倒排索引 传统数据库为特定列增加一个索引，例如B-Tree索引来加速检索。Elasticsearch和Lucene使用一种叫做倒排索引(inverted index)的数据结构来达到相同目的。

```
PUT /megacorp/employee/1
{
    "first_name" : "John",
    "last_name" :  "Smith",
    "age" :        25,
    "about" :      "I love to go rock climbing",
    "interests": [ "sports", "music" ]
}
```
我们看到path:/megacorp/employee/1包含三部分信息：    

|名字     |	说明|    
|:-----  |:-----|  
|megacorp|	索引名|    
|employee|	类型名|    
|1       |	这个员工的ID|
### ports
9200作为Http协议，主要用于外部通讯a
9300作为Tcp协议，jar之间就是通过tcp协议通讯
ES集群之间是通过9300进行通讯

## books
https://wiki.jikexueyuan.com/project/elasticsearch-definitive-guide-cn/010_Intro/25_Tutorial_Indexing.html
https://learnku.com/docs/elasticsearch73/7.3/data-in-documents-and-indices/6446
https://es.xiaoleilu.com/index.html  