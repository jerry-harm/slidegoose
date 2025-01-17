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
	if file_count == 0 {
		c.JSON(400, gin.H{"status": "nothing added"})
		return
	}
	c.JSON(201, gin.H{"status": "200", "created": file_count})
}
