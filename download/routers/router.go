package routers

import (
	v1 "deepgo/download/api/v1"
	"deepgo/download/middleware"
	"deepgo/download/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	r.Use(middleware.Cors())

	apiv1 := r.Group("/api/v1")
	{
		//获取软件列表
		apiv1.GET("/software", v1.GetSoftware)

	}

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "static/html")
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	r.Use(middleware.Cors())
	return r
}
