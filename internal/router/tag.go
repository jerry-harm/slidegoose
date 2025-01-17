package router

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"slidegoose/internal/database"
)

func AddTag(c *gin.Context) {
	type addTagForm struct {
		Name        string `binding:"required"`
		Description string
	}

	var form addTagForm
	if err := c.Bind(&form); err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
		return
	}

	database.AddTag(form.Name, form.Description)
	c.JSON(200, gin.H{"status": "OK"})
}

func setTag(c *gin.Context, fn func(uint, uint) error) {
	type setTagForm struct {
		Tag uint
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
		return
	}

	var form setTagForm
	if err := c.Bind(&form); err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
		return
	}

	if err := fn(uint(id), form.Tag); err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "OK"})
}

func SetVideoTag(c *gin.Context) {
	setTag(c, database.SetVideoTag)
}

func SetClipTag(c *gin.Context) {
	setTag(c, database.SetClipTag)
}

func SetPictureTag(c *gin.Context) {
	setTag(c, database.SetPictureTag)
}

func SetAudioTag(c *gin.Context) {
	setTag(c, database.SetAudioTag)
}
