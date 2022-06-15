package config

type Config struct {
	Redis Redis `json:"redis" yaml:"redis" mapstructure:"redis"`
}
