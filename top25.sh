#!/bin/bash

echo "## TOP 25 Client IPs"
docker-compose exec ch-server clickhouse-client --query "SELECT count(), clientip FROM burst GROUP BY clientip ORDER BY count() DESC LIMIT 25"
echo
echo "## TOP 25 Referers"
docker-compose exec ch-server clickhouse-client --query "SELECT count(), referer FROM burst GROUP BY referer ORDER BY count() DESC LIMIT 25"
