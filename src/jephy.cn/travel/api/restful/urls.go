package restful

import (
	"github.com/gin-gonic/gin"
	"jephy.cn/travel/api/restful/article"
	"jephy.cn/travel/api/restful/class"
)

func RegisterURLs(r gin.IRoutes) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})

	article.RegisterURLs(r)
	class.RegisterURLs(r)
}
