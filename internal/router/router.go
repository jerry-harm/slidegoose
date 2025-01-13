package router

import (
	"log"
	"strconv"

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
	return router
}

func AddFile(c *gin.Context) {
	type filesForm struct {
		Files []string
	}

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

func AddTag(c *gin.Context) {
	type tagForm struct {
		Name        string `binding:"required"`
		Description string
	}

	var form tagForm
	if err := c.Bind(&form); err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
		return
	}

	database.AddTag(form.Name, form.Description)
	c.JSON(200, gin.H{"status": "OK"})
}

func SetVideoTag(c *gin.Context) {
	type tagForm struct {
		Tag uint
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panicln(err)
		c.JSON(400, gin.H{"status": err.Error()})
		return
	}

	var form tagForm
	if err := c.Bind(&form); err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
		return
	}

	if err := database.SetVideoTag(uint(id), form.Tag); err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "OK"})
}
