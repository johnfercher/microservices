# Execute Application

### 1 - Execute Entire Solution
```
$ bash run.sh
```

### 2 - Create Kafka Topic
```
$ create_topic.sh
$ Kafka Topic Name: topic-user-events
$ Creating kafka-topic topic-user-events.
```

# Setup Application

### 1 - Setup Elasticsearch Folder
```
cd /usr/share
mkdir elasticsearch/data/nodes
chmod 777 -R /elasticsearch/data
```

### 2 - Build UserDb Docker Image
```
cd userdb
bash build.sh
```

### 3 - Execute Entire Solution
```
$ bash run.sh
```

### 4 - Create Kafka Topic
```
$ create_topic.sh
$ Kafka Topic Name: topic-user-events
$ Creating kafka-topic topic-user-events.
```

# Addresses
* Kibana
    * http://localhost:5601
* Elasticsearch
    * http://localhost:9200
* Zookeeper
    * http://localhost:2181
* Kafka1
    * http://localhost:9092
* MySQL
    * http://localhost:3306

# Reference

### Kafka
* [Apache Kafka — Aprendendo na prática][kafka_evandro]
* [Kafka Base Examples][kafka_examples]

### Logstach, Elasticsearch and Kibana
* [Configurando o Elasticsearch e Kibana no Docker][log_docker]

[kafka_evandro]: https://medium.com/trainingcenter/apache-kafka-codifica%C3%A7%C3%A3o-na-pratica-9c6a4142a08f
[kafka_examples]: https://github.com/confluentinc/cp-docker-images/tree/5.3.3-post/examples
[log_docker]: https://medium.com/@hgmauri/configurando-o-elasticsearch-e-kibana-no-docker-3f4679eb5feb