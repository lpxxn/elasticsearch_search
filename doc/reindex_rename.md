## Change the mapping of an existing fieldedit
Except for supported mapping parameters, you can’t change the mapping or field type of an existing field. Changing an existing field could invalidate data that’s already indexed.
If you need to change the mapping of a field, create a new index with the correct mapping and reindex your data into that index.
To see how this works, try the following example.
Use the create index API to create the users index with the user_id field with the long field type.
```
curl -X PUT "localhost:9200/users?pretty" -H 'Content-Type: application/json' -d'
{
  "mappings" : {
    "properties": {
      "user_id": {
        "type": "long"
      }
    }
  }
}
'


{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "users"
}
```

Use the index API to index several documents with user_id field values.

```
curl -X POST "localhost:9200/users/_doc?refresh=wait_for&pretty" -H 'Content-Type: application/json' -d'
{
    "user_id" : 12345
}
'
curl -X POST "localhost:9200/users/_doc?refresh=wait_for&pretty" -H 'Content-Type: application/json' -d'
{
    "user_id" : 12346
}
'

{
  "_index" : "users",
  "_type" : "_doc",
  "_id" : "syZvf3ABNd7cO_9kKAJL",
  "_version" : 1,
  "result" : "created",
  "_shards" : {
    "total" : 2,
    "successful" : 1,
    "failed" : 0
  },
  "_seq_no" : 0,
  "_primary_term" : 1
}
{
  "_index" : "users",
  "_type" : "_doc",
  "_id" : "tCZvf3ABNd7cO_9kKgJP",
  "_version" : 1,
  "result" : "created",
  "_shards" : {
    "total" : 2,
    "successful" : 1,
    "failed" : 0
  },
  "_seq_no" : 1,
  "_primary_term" : 1
}
```

To change the user_id field to the keyword field type, use the create index API to create the new_users index with the correct mapping.

```
curl -X PUT "localhost:9200/new_users?pretty" -H 'Content-Type: application/json' -d'
{
  "mappings" : {
    "properties": {
      "user_id": {
        "type": "keyword"
      }
    }
  }
}
'

```

Use the reindex API to copy documents from the users index to the new_users index.


```
curl -X POST "localhost:9200/_reindex?pretty" -H 'Content-Type: application/json' -d'
{
  "source": {
    "index": "users"
  },
  "dest": {
    "index": "new_users"
  }
}
'

```

search 
```
curl "http://localhost:9200/users/_search?pretty"

curl "http://localhost:9200/new_users/_search?pretty"

```


## Rename a field
Renaming a field would invalidate data already indexed under the old field name. Instead, add an alias field to create an alternate field name.
For example, use the create index API to create an index with the user_identifier field.

```
curl -X PUT "localhost:9200/my_index?pretty" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "properties": {
      "user_identifier": {
        "type": "keyword"
      }
    }
  }
}
'


```
if the index exists
```
curl -XPUT "http://localhost:9200/my_index/_mapping/?pretty" -H 'Content-Type: application/json' -d '
{
    "properties": {
      "user_identifier": {
        "type": "keyword"
      }
    }
}
'
```

Use the put mapping API to add the user_id field alias for the existing user_identifier field.

```
curl -X PUT "localhost:9200/my_index/_mapping?pretty" -H 'Content-Type: application/json' -d'
{
  "properties": {
    "user_id": {
      "type": "alias",
      "path": "user_identifier"
    }
  }
}
'

```
Use the get mapping API to verify your changes.

```
{
  "my_index" : {
    "mappings" : {
      "properties" : {
        "user_id" : {
          "type" : "alias",
          "path" : "user_identifier"
        },
        "user_identifier" : {
          "type" : "keyword"
        }
      }
    }
  }
}
```
