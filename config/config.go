package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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

	execPath, err := os.Executable()
    if err != nil {
        fmt.Println(err)
    }

	if !filepath.IsAbs(viper.GetString("web")){
		viper.Set("web",
		filepath.Join(filepath.Dir(execPath),
			filepath.Dir(viper.GetString("web"))))
	}

	log.Println("web page in",viper.GetString("web"))

	return *viper.GetViper()
}
