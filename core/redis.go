package core

import (
	"fmt"
	"github.com/go-redis/redis"
	"wTools/global"
)

func InitRedis() {
	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:       redisCfg.Addr,
		Password:   redisCfg.Password,
		DB:         redisCfg.Db,
		MaxRetries: redisCfg.MaxRetries,
		PoolSize:   redisCfg.PoolSize,
	})
	fmt.Println(redisCfg)
	p, err := client.Ping().Result()
	if err != nil {
		fmt.Println(p, err)
	} else {
		global.Redis = client
	}
}
