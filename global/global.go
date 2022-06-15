package global

import (
	"github.com/go-redis/redis"
	"wTools/config"
)

var (
	Config config.Config
	Redis  *redis.Client
)
