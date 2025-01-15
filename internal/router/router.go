package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"slidgoose/internal/database"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/", viper.GetString("web"))
	router.POST("/file", AddFile)
	router.POST("/tag", AddTag)
	router.POST("/video/:id/tag", SetVideoTag)
	router.POST("/clip/:id/tag", SetClipTag)
	router.POST("/audio/:id/tag", SetAudioTag)
	router.POST("/picture/:id/tag", SetPictureTag)
	return router
}

type filesForm struct {
	Files []string
}

func AddFile(c *gin.Context) {

	var form filesForm
	if err := c.Bind(&form); err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
		return
	}

	file_count := 0
	for _, v := range form.Files {
		log.Println("adding", v)
		file_count += database.AddDir(v)
	}

	c.JSON(200, gin.H{"status": "200", "OK": file_count})
}
