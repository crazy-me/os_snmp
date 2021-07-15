package initialize

import (
	"fmt"
	"github.com/crazy-me/os_snmp/utils/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/toolkits/pkg/file"
)

const configFile = "app.yaml"

func LoadConfInit(conf string) {
	if conf == "" || !file.IsExist(conf) {
		conf = configFile
	}
	v := viper.New()
	v.SetConfigFile(conf)
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
