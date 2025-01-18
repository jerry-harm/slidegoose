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
