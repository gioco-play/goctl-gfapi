package config

import {{.authImport}}

type Config struct {
	rest.RestConf
	{{.auth}}
	{{.jwtTrans}}
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
        DBPoolMin         int
        DBPoolMax         int
        DBConnMaxLifetime int
        DBDebugLevel      string
    }
    RedisCache struct {
        RedisSentinelNode string
        RedisMasterName   string
        RedisDB           int
    }
}
