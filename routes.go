package main

import (
	"github.com/gin-gonic/gin"
	"go+vue/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Reginster)
	r.POST("/api/auth/login", controller.Login)
	return r
}
