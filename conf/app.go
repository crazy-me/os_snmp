package conf

type App struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`

	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`

	Consul Consul `mapstructure:"consul" json:"consul" yaml:"consul"`
}
