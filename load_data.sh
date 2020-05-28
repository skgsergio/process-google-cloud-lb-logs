#!/bin/bash

# Load data
docker-compose exec ch-server clickhouse-client --query "CREATE TABLE requests (timestamp DateTime, clientip String, referer String) ENGINE = Log;"
docker-compose exec -T ch-server clickhouse-client --query="INSERT INTO requests FORMAT CSVWithNames" < requests.csv

# Aggregate IPs per timestamp
docker-compose exec ch-server clickhouse-client --query="CREATE TABLE requests_ips ENGINE=Log AS SELECT timestamp, clientip, count() AS reqs FROM requests GROUP BY timestamp, clientip ORDER BY timestamp"

# Aggregate referer per timestamp
docker-compose exec ch-server clickhouse-client --query="CREATE TABLE requests_referer ENGINE=Log AS SELECT timestamp, referer, count() AS reqs FROM burst GROUP BY timestamp, referer ORDER BY timestamp"
