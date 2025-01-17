package router

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"slidegoose/web"

	_ "slidegoose/docs"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	// static file
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/web")
	})
	staticFS, _ := fs.Sub(web.WebFS, "dist")
	router.StaticFileFS("/favicon.ico", "/favicon.ico", http.FS(staticFS))
	router.StaticFS("/web", http.FS(staticFS))

	if gin.Mode() != gin.ReleaseMode {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

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
