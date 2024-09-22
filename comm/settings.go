package comm

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var CfgLoader *viper.Viper

func InitViperCfg() (err error) {
	CfgLoader = viper.New()
	CfgLoader.SetConfigFile("./conf/conf.yaml")
	CfgLoader.AddConfigPath(".")
	err = CfgLoader.ReadInConfig()
	if err != nil {
		fmt.Println("load config failed!")
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("load config success!")
	CfgLoader.WatchConfig()
	CfgLoader.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed!")
	})

	return nil
}
