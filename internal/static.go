package internal

import (
	"net/http"

	"github.com/spf13/viper"
)

func IndexPage() http.Handler {
	return http.FileServer(http.Dir(viper.GetString("web")))
}
