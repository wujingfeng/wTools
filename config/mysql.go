package config

type Mysql struct {
	Host         string `json:"host" mapstructure:"host" yaml:"host"`
	Port         string `json:"port" mapstructure:"port" yaml:"port"`
	Username     string `json:"username" mapstructure:"username" yaml:"username"`
	Password     string `json:"password" mapstructure:"password" yaml:"password"`
	Db           string `json:"db" mapstructure:"db" yaml:"db"`
	Config       string `json:"config" mapstructure:"config" yaml:"config"`                   //
	MaxIdleConns int    `json:"MaxIdleConns" mapstructure:"MaxIdleConns" yaml:"MaxIdleConns"` // Default
	MaxOpenConns int    `json:"MaxOpenConns" mapstructure:"MaxOpenConns" yaml:"MaxOpenConns"`
}

func (m *Mysql) DSN() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Db + "?" + m.Config
}
