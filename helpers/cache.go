package helpers

import "github.com/go-redis/redis"

type cache struct {
}

var RedisClient *redis.Client
