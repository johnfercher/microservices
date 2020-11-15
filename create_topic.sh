#!/bin/bash

read -p 'Kafka Topic Name: ' kafkaTopicName
echo
echo Creating kafka-topic $kafkaTopicName.

docker-compose exec kafka  \
kafka-topics --create --topic $kafkaTopicName --partitions 1 --replication-factor 1 --if-not-exists --zookeeper zookeeper:2181