package config

import {{.authImport}}
import {{.serviceImportStr}}

type Config struct {
	rest.RestConf
	{{.auth}}
	{{.jwtTrans}}
    service.ServiceConf
    Consul struct {
        Target string
    }
    Postgres struct {
        Host              string
        Port              int
        SlavePort         int
        UserName          string
        Password          string
        DBName            string
        DBTimezone        string
        DBPoolSize        int
        DBConnMaxLifetime int
        DBDebugLevel      string
    }
    RedisCache struct {
        RedisSentinelNode string
        RedisMasterName   string
        RedisDB           int
    }
}
