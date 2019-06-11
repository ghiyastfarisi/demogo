package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	// Env is environment
	Env = os.Getenv("APP_ENV")
	// Element is element of configuration data
	Element Config
)

// Init reading configuration file
func Init() {
	if Env == "" {
		Env = "development"
	}
	viper.SetConfigName(fmt.Sprintf("%s.conf", Env))
	viper.AddConfigPath("/etc/demogo/configs/")
	viper.AddConfigPath("./configs/")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("[ERR] Fatal error config file: %s", err))
	}
	if err := viper.Unmarshal(&Element); err != nil {
		panic(fmt.Errorf("[ERR] Cannot unmarshal config: %s", err))
	}
}
