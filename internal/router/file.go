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
// @Tags         system
// @Produce      json
// @Router       /file [post]
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
