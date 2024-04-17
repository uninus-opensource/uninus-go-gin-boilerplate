package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/feature/auth"
)

func AuthRoute(r *gin.Engine, h auth.AuthHandlerInterface) {
	authGroup := r.Group("/v1")
	{
		authGroup.POST("/auth/login", h.Login)
		authGroup.POST("/auth/register", h.Register)
	}
}
