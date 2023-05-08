package routes

import (
	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {

	v01 := router.Group("/v1")
	{
		v01.GET("/user")
		v01.GET("/user/:id")
		v01.POST("/user")
		v01.DELETE("/user/:id")
		v01.PUT("/user/:id")
	}

	return v01
}
