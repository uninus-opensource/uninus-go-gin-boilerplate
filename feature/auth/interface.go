package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/entities"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/feature/auth/dto"
)

type AuthSvcInterface interface {
	Register(dto *dto.TypeRegisterRequest) (*entities.MstUser, error)
	Login(dto *dto.TypeLoginRequest) (*entities.MstUser, string, error)
}

type AuthHandlerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}
