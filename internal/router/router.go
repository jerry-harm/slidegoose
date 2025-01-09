package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/",viper.GetString("web"))
	return router
}

