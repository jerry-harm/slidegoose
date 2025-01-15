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

	api := router.Group("api")
	{
		file := api.Group("file")
		{
			file.POST("", AddFile)
		}
		tag := api.Group("tag")
		{
			tag.POST("", AddTag)
		}
		video := api.Group("video")
		{
			video.POST("/:id/tag", SetVideoTag)
		}
		clip := api.Group("clip")
		{
			clip.POST("/:id/tag", SetClipTag)
		}
		picture := api.Group("picture")
		{
			picture.POST("/:id/tag", SetPictureTag)
		}
		audio := api.Group("audio")
		{
			audio.POST("/:id/tag", SetAudioTag)
		}
	}

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
