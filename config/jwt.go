package config

type JWT struct {
	SigningKey  string `json:"signing_key" mapstructure:"signingKey" yaml:"signingKey"`
	ExpiresTime int64  `json:"expires_time" mapstructure:"expiresTime" yaml:"expiresTime"`
	BufferTime  int64  `json:"buffer_time" mapstructure:"bufferTime" yaml:"bufferTime"`
}
