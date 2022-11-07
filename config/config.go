package config

type Config struct {
	Redis Redis `json:"redis" yaml:"redis" mapstructure:"redis"`
	JWT   JWT   `json:"jwt" yaml:"jwt" mapstructure:"jwt"`
	Mysql Mysql `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
}
