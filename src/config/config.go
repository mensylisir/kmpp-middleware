package config

import (
	"github.com/mensylisir/kmpp-middleware/src/util/logx"
	"github.com/spf13/viper"
)

var (
	C = new(Config)
)

type Config struct {
	RunMode string
	Log     logx.Config
}

func Init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	//viper.AddConfigPath("C:\\Users\\mensyli1\\work\\workspace\\middleware\\conf")
	viper.AddConfigPath("/etc/middleware")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}
