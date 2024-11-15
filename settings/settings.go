package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed with %s\n", err)
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed")

		// 重新读取配置文件
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("viper.ReadInConfig() after change failed with %s\n", err)
		} else {
			fmt.Println("Successfully re-read config after change")
		}
	})
	return
}
