#!/bin/bash

read -p 'Container ID: ' containerId
echo
echo Loading data into $containerId.

cat sql-scripts/create_tables.sql | docker exec -i $containerId mysql -u root --password=supersecret UserDb
cat sql-scripts/insert_data.sql | docker exec -i $containerId mysql -u root --password=supersecret UserDb