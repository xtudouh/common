package redis

import (
    "gopkg.in/redis.v2"
    "xtudouh/common/conf"
    "xtudouh/common/log"
    "time"
)

/*******************EXAMPLE**********************
* Reference: https://godoc.org/gopkg.in/redis.v2
*
************************************************/

var (
    RedisClient *redis.Client
    l = log.NewLogger()
)

func Init() {
    l.Debug("redis module start initializing.")
    server := conf.String("redis", "REDIS_SERVER", ":6379")

    RedisClient = redis.NewClient(&redis.Options{
        Network: "tcp",
        Addr:    server,

        DialTimeout:  5 * time.Second,
        ReadTimeout:  time.Second,
        WriteTimeout: time.Second,

        PoolSize:    conf.Int("redis", "POOL_SIZE", 50),
        IdleTimeout: time.Second * time.Duration(conf.Int64("redis", "IDEL_TIMEOUT", 600)),
    })

    l.Debug("redis module initialize successfully.")
}
