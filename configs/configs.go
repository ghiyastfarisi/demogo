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
	filepath := "development"
	if Env != "development" && Env != "" {
		filepath = Env
	}
	viper.SetConfigName("conf")
	viper.AddConfigPath(fmt.Sprintf("/etc/demogo/config/%s/", filepath))
	viper.AddConfigPath(fmt.Sprintf(".%s/", filepath))
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("[ERR] Fatal error config file: %s", err))
	}
	if err := viper.Unmarshal(&Element); err != nil {
		panic(fmt.Errorf("[ERR] Cannot unmarshal config: %s", err))
	}
}
