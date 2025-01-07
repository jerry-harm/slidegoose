package query

import (
	"net/http"

	"github.com/spf13/viper"
)

func GetMux() *http.ServeMux{
	mux := http.NewServeMux()
	mux.Handle("/",indexpage())
	return mux
}


func indexpage() http.Handler {
	return http.FileServer(http.Dir(viper.GetString("web")))
}
