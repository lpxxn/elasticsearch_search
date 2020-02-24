
## docker
https://www.elastic.co/guide/en/elasticsearch/reference/7.6/docker.html


```
docker pull elasticsearch:7.6.0
```
### single node
```
docker run --rm -v /Users/lipeng/temp/share/elasticsearch/data:/usr/share/elasticsearch/data -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.6.0
```

#### dev model
```
docker network create somenetwork
```


```
docker run -d --name elasticsearch --net somenetwork -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:tag
```

