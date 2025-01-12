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
	return router
}

func AddFile(c *gin.Context) {
	type FilesForm struct {
		Files []string `form:"files" json:"files" xml:"files" binding:"required"`
	}

	var form FilesForm
	if err := c.Bind(&form); err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
		return
	}

	file_count := 0
	for _, v := range form.Files {
		log.Println("adding", v)
		file_count += database.AddDir(v)
	}

	c.JSON(200, gin.H{"status": "added", "added": file_count})
}
