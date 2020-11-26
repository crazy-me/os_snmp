package conf

type Zap struct {
	LogPrefix     string `mapstructure:"log-prefix" json:"logPrefix" yaml:"log-prefix"`
	LogPath       string `mapstructure:"log-path" json:"logPath"  yaml:"log-path"`
	LogMaxSize    int    `mapstructure:"log-max-size" json:"logMaxSize"  yaml:"log-max-size"`
	LogMaxBackups int    `mapstructure:"log-max-backups" json:"logMaxBackups"  yaml:"log-max-backups"`
	LogMaxAge     int    `mapstructure:"log-max-age" json:"logMaxAge"  yaml:"log-max-age"`
	LogCompress   bool   `mapstructure:"log-compress" json:"logCompress"  yaml:"log-compress"`
}
