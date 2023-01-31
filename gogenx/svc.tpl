package svc

import (
    "fmt"
	{{.configImport}}

	"github.com/go-redis/redis/v8"
    "github.com/neccoys/go-driver/postgrex"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "strings"
)

type ServiceContext struct {
	Config {{.config}}
	RedisClient  *redis.Client
	BoDB         *gorm.DB
	{{.middleware}}
}

func NewServiceContext(c {{.config}}) *ServiceContext {
	// Redis
	redisClient := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    c.RedisCache.RedisMasterName,
		SentinelAddrs: strings.Split(c.RedisCache.RedisSentinelNode, ";"),
		DB:            c.RedisCache.RedisDB,
	})

	// DB
	db, err := postgrex.New(c.Postgres.Host, fmt.Sprintf("%d", c.Postgres.Port), c.Postgres.UserName,
        c.Postgres.Password, c.Postgres.DBName).
        SetTimeZone(c.Postgres.DBTimezone).
        SetLogger(logger.Default.LogMode(postgrex.Level(c.Postgres.DBDebugLevel))).
        Connect(postgrex.Pool(c.Postgres.DBPoolMin, c.Postgres.DBPoolMax, c.Postgres.DBConnMaxLifetime))

	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
        RedisClient:  redisClient,
        BoDB:         db,
		{{.middlewareAssignment}}
	}
}
