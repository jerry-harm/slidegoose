package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"

	"slidgoose/config"
	"slidgoose/internal/query"
)

func main() {

	config.ReadConfig()

	mux := query.GetMux()
	log.Printf("open on address http://%s:%s", viper.GetString("host"), viper.GetString("port"))
	http.ListenAndServe(viper.GetString("ip")+":"+viper.GetString("port"), mux)

}
