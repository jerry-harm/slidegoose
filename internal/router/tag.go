package router

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"slidegoose/internal/database"
)

type addTagForm struct {
	Name        string `binding:"required"`
	Description string
}

// AddTagDoc
// @Summary      add Tag
// @Description  return status
// @param       name body    addTagForm  true  "tag"
// @Tags         tag
// @Produce      json
// @Router       /tag [post]
func AddTag(c *gin.Context) {

	var form addTagForm
	if err := c.Bind(&form); err != nil {
		c.JSON(400, gin.H{"status": err.Error()})
		return
	}

	database.AddTag(form.Name, form.Description)
	c.JSON(201, gin.H{"status": "OK"})
}

type setTagForm struct {
	Tag uint
}

func setTag(c *gin.Context, fn func(uint, uint) error) {

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
	c.JSON(201, gin.H{"status": "OK"})
}

// SetVideoTagDoc
// @Summary      set video Tag
// @Description  return status
// @param       name body    setTagForm  true  "tag"
// @param       videoId path    uint  true  "video id"
// @Tags         tag
// @Produce      json
// @Router       /video/{videoId}/tag [post]
func SetVideoTag(c *gin.Context) {
	setTag(c, database.SetVideoTag)
}

// SetClipTagDoc
// @Summary      set clip Tag
// @Description  return status
// @param       name body    setTagForm  true  "tag"
// @param       clipId path    uint  true  "clip id"
// @Tags         tag
// @Produce      json
// @Router       /clip/{clipId}/tag [post]
func SetClipTag(c *gin.Context) {
	setTag(c, database.SetClipTag)
}

// SetPictureTagDoc
// @Summary      set picture Tag
// @Description  return status
// @param       name body    setTagForm  true  "tag"
// @param       pictureId path    uint  true  "picture id"
// @Tags         tag
// @Produce      json
// @Router       /picture/{pictureId}/tag [post]
func SetPictureTag(c *gin.Context) {
	setTag(c, database.SetPictureTag)
}

// SetAudioTagDoc
// @Summary      set audio Tag
// @Description  return status
// @param       name body    setTagForm  true  "tag"
// @param       audioId path    uint  true  "audio id"
// @Tags         tag
// @Produce      json
// @Router       /audio/{audioId}/tag [post]
func SetAudioTag(c *gin.Context) {
	setTag(c, database.SetAudioTag)
}
