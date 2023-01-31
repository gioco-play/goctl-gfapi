Name: ${APP_NAME}
Host: ${APP_HOST}
Port: ${APP_PORT}
Timeout: ${TIMEOUT}

Log:
  Mode: ${LOG_MODE}

Telemetry:
  Name: ${APP_NAME}
  Endpoint: ${TRACE_ENDPOINT}
  Sampler: 1.0
  Batcher: jaeger

Postgres:
  Host: ${DB_HOST}
  Port: ${DB_PORT}
  SlavePort: ${DB_SLAVE_PORT}
  UserName: ${DB_USERNAME}
  Password: ${DB_PASSWORD}
  DBName: ${DB_DATABASE}
  DBPoolMin: ${DB_POOL_MIN}
  DBPoolMax: ${DB_POOL_MAX}
  DBTimezone: ${DB_TIMEZONE}
  DBConnMaxLifetime: ${DB_CONN_MAX_LIFETIME}
  DBDebugLevel: ${DB_DEBUG_LEVEL}

RedisCache:
  RedisSentinelNode: ${REDIS_SENTINEL_NODE}
  RedisMasterName: ${REDIS_MASTER_NAME}
  RedisDB: ${REDIS_DB}

Consul:
  Target: ${CONSUL_HOST}

Prometheus:
  Host: ${PROMETHEUS_HOST}
  Port: ${PROMETHEUS_PORT}
  Path: ${PROMETHEUS_PATH}