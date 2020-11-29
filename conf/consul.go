package conf

type Consul struct {
	Address    string `mapstructure:"address" json:"address" yaml:"address"`
	Interval   int    `mapstructure:"interval" json:"interval" yaml:"interval"`
	ExpireTime int    `mapstructure:"expire-time" json:"expire-time" yaml:"expire-time"`
}
