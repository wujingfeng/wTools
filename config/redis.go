package config

import "time"

type Redis struct {
	Addr       string        `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password   string        `mapstructure:"password" json:"password" yaml:"password"`
	Db         int           `mapstructure:"db" json:"db" yaml:"db"`
	MaxRetries int           `mapstructure:"maxRetries" json:"maxRetries" yaml:"maxRetries"`
	PoolSize   int           `mapstructure:"poolSize" json:"poolSize" yaml:"poolSize"`
	Prefix     string        `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	TTL        time.Duration `mapstructure:"ttl" json:"ttl" yaml:"ttl"`
}
