# Google LB Logs parser

Quick'n'dirty environment for processing Google Cloud LB logs.

## Components

- `parse.go`: Parses Google Cloud LB logs that contains one request per line in
  JSON format.

- `docker-compose.yml`:
  - `grafana`: Grafana exposed in port 3000.
    - `grafana-clickhouse-ds.yml`: Grafana's ClickHouse datasource configuration.
  - `ch-server`: ClickHouse server for loading the requests and doing all the
    queries and agregations you need.

- `load_data.sh`: Script for loading `parse.go` results into ClickHouse in
  `requests` table. It also creates two aggregated tables, one per timestamp and
  IPs named `requests_ips` and other one per timestamp and referer named
  `requests_referer`.

- `top25.sh`: Script for querying the top 25 requests and referers.

## Usage

```bash
go run parse.go /path/to/logs/*.json
./load_data.sh
./top25.sh
```

For running the ClickHouse CLI just run: `docker-compose exec ch-server clickhouse-client`

Grafana might be also helpful for plotting requests over time and so using it's
explore mode, just open http://localhost:3000/ and use Explore, it's pre-configured.
