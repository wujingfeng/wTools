package core

import (
	"fmt"
	"github.com/go-redis/redis"
	"wTools/global"
)

func InitRedis() {
	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	p, err := client.Ping().Result()
	if err != nil {
		fmt.Println(p, err)
	} else {
		global.Redis = client
	}
}
