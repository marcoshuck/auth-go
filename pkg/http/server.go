package http

import (
	"github.com/gin-gonic/gin"
	"github.com/marcoshuck/auth-go/pkg/controller"
)

func NewServer() *gin.Engine {
	return gin.Default()
}

func Register(router *gin.Engine, auth controller.AuthController) *gin.Engine {
	group := router.Group("/auth")
	{
		group.POST("/login", auth.Login)
		group.POST("/register", auth.Register)
	}
	return router
}
