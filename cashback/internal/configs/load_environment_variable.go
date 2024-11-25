package configs

import (
	"path/filepath"

	viper "github.com/spf13/viper"
)

func Setup() {
	dir, _ := filepath.Abs("./")

	viper.AddConfigPath(dir)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.AutomaticEnv()
}
