package svc

import (
	{{.configImport}}
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
	db, err := postgrez.New(c.Postgres.Host, fmt.Sprintf("%d", c.Postgres.Port), c.Postgres.UserName,
		c.Postgres.Password, c.Postgres.DBName).
		SetTimeZone(c.Postgres.DBTimezone).
		SetLogger(logrusz.New().SetLevel(c.Postgres.DBDebugLevel).Writer()).
		Connect(postgrez.Pool(c.Postgres.DBPoolSize, c.Postgres.DBPoolSize, c.Postgres.DBConnMaxLifetime))

	if err != nil {
		panic(err)
	}

    mongo, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.MongoDB.URI))
    if err != nil {
        panic(err)
    }

	// Tracer
	ztrace.StartAgent(ztrace.Config{
		Name:     c.Telemetry.Name,
		Endpoint: c.Telemetry.Endpoint,
		Batcher:  c.Telemetry.Batcher,
		Sampler:  c.Telemetry.Sampler,
	})

	return &ServiceContext{
		Config: c,
        RedisClient:  redisClient,
        BoDB:         db,
		{{.middlewareAssignment}}
	}
}
