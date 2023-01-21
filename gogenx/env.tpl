APP_NAME={{.serviceName}}
APP_MODE=debug
APP_HOST={{.host}}
APP_PORT={{.port}}
TIMEOUT=20000

DB_HOST=127.0.0.1
DB_PORT=5432
DB_SLAVE_PORT=5432
DB_DATABASE=test_db
DB_USERNAME=postgres
DB_PASSWORD=
DB_POOL_SIZE=10
DB_TIMEZONE=UTC
DB_CONN_MAX_LIFETIME=180
DB_DEBUG_LEVEL=error

REDIS_SENTINEL_NODE=192.168.30.38:26379;192.168.30.39:26379;192.168.30.40:26379
REDIS_MASTER_NAME=mymaster
REDIS_DB=1

CONSUL_HOST=192.168.10.93:8500
CONSUL_TTL=20

TRACE_ENDPOINT=http://172.16.204.124:14268/api/traces

LOG_MODE=console

PROMETHEUS_HOST=0.0.0.0
PROMETHEUS_PORT=9091
PROMETHEUS_PATH=/metrics