package config

import (
	"log"

	"github.com/spf13/viper"
)

func ReadConfig() viper.Viper {
	viper.SetDefault("port", "5444")
	viper.SetDefault("host", "127.0.0.1")
	viper.SetDefault("database", "data.db")
	viper.SetDefault("web", "./web/")

	viper.SetConfigName("config.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	return *viper.GetViper()
}
