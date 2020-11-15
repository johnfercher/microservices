#!/bin/bash

docker run --volume=$HOME/datadir:/var/lib/mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=supersecret -e MYSQL_DATABASE=UserDb -e MYSQL_USER=AdminUser -e MYSQL_PASSWORD=AdminPassword user-db