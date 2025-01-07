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

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		// create default config file
		viper.SafeWriteConfig()
	}


	log.Println("web page in",viper.GetString("web"))

	return *viper.GetViper()
}