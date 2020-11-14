#!/bin/bash

docker-compose exec kafka  \
kafka-topics --create --topic kafka-topic-test --partitions 1 --replication-factor 1 --if-not-exists --zookeeper zookeeper:2181