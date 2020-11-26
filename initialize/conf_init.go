package initialize

import (
	"fmt"
	"github.com/crazy-me/os_snmp/utils/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const configFile = "app.yaml"

func init() {
	v := viper.New()
	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Configuration file failed to load: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		err1 := v.Unmarshal(&global.APP)
		if err1 != nil {
			fmt.Println(err)
		}

	})

	err2 := v.Unmarshal(&global.APP)
	if err2 != nil {
		fmt.Println(err2)
	}
}
