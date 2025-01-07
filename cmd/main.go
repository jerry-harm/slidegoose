package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"

	"slidgoose/config"
)

func main() {

	config.ReadConfig()

	http.Handle("/", http.FileServer(http.Dir(viper.GetString("web"))))
	log.Printf("open on address http://%s:%s", viper.GetString("host"), viper.GetString("port"))
	http.ListenAndServe(
		viper.GetString("ip")+":"+viper.GetString("port"), nil)

}
