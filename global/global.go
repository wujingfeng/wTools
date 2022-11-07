package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"wTools/config"
)

var (
	Config config.Config
	Redis  *redis.Client
	Viper  *viper.Viper
	Mysql  *gorm.DB
)
