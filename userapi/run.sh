#!/bin/bash

read -p 'Container ID: ' containerId
echo
echo Connecting to network of $containerId.

docker run -it --rm --network container:$containerId -p 8080:8080 user-api