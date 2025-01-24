package router

import (
	"log"

	"github.com/gin-gonic/gin"

	"slidegoose/internal/database"
)

type filesForm struct {
	Files []string
}

// AddFileDoc
// @Summary      add file or files
// @Description  return number of added files
// @param       name body    filesForm  true  "files"
// @Tags         file
// @Produce      json
// @Router       /file [post]
// @Success 	201	{object} JSONResult{data=int}
// @Failure 	400 {object} JSONResult
func AddFile(c *gin.Context) {

	var form filesForm
	if err := c.Bind(&form); err != nil {
		c.JSON(400, JSONResult{Code: 400, Message: err.Error()})
		return
	}

	file_count := 0
	for _, v := range form.Files {
		log.Println("adding", v)
		file_count += database.AddDir(v)
	}
	if file_count == 0 {
		c.JSON(400, JSONResult{Code: 400, Message: "nothing added"})
		return
	}
	c.JSON(201, JSONResult{Code: 201, Data: file_count})
}

// GetVideoInfo
// @Description  return video info
// @param       id  path    int  true  "id"
// @Tags         file
// @Produce      json
// @Router       /video/{id}/info [get]
// @Success 	200	{object} JSONResult
// @Failure 	404
func GetVideoInfo(c *gin.Context) {
	id := c.Param("id")
	var video database.Video
	if err := database.DB.First(&video, id).Error; err != nil {
		c.Status(404)
		return
	}
	c.JSON(200, JSONResult{Code: 200, Data: video})
}

// GetPictureInfo
// @Description  return picture info
// @param       id  path    int  true  "id"
// @Tags         file
// @Produce      json
// @Router       /picture/{id}/info [get]
// @Success 	200	{object} JSONResult
// @Failure 	404
func GetPictureInfo(c *gin.Context) {
	id := c.Param("id")
	var picture database.Picture
	if err := database.DB.First(&picture, id).Error; err != nil {
		c.Status(404)
		return
	}
	c.JSON(200, JSONResult{Code: 200, Data: picture})
}
