package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"slidgoose/config"
	"slidgoose/internal/query"
)

func main() {
	// chdir to right place
	execPath, err := os.Executable()
	if err != nil {log.Fatalln(err)}
	execPath = filepath.Dir(execPath)
	os.Chdir(execPath)


	config.ReadConfig()

	mux := query.GetMux()
	log.Printf("open on address http://%s:%s", viper.GetString("host"), viper.GetString("port"))
	http.ListenAndServe(viper.GetString("ip")+":"+viper.GetString("port"), mux)

}
