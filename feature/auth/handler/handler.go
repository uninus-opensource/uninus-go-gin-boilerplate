package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/feature/auth"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/feature/auth/dto"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/helper/response"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/helper/validator"
)

type authHandler struct {
	authSvc auth.AuthSvcInterface
}

func NewAuthHandler(authSvc auth.AuthSvcInterface) auth.AuthHandlerInterface {
	return &authHandler{
		authSvc: authSvc,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var req dto.TypeRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
		return
	}

	if err := validator.ValidateStruct(req); err != nil {
		response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
		return
	}

	_, err := h.authSvc.Register(&req)
	if err != nil {
		response.SendStatusBadRequest(c, err.Error())
		return
	}
	response.SendStatusOkResponse(c, "register is succesfully")
}

func (h *authHandler) Login(c *gin.Context) {
	var req dto.TypeLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
		return
	}

	if err := validator.ValidateStruct(req); err != nil {
		response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
		return
	}

	user, accessToken, err := h.authSvc.Login(&req)
	if err != nil {
		response.SendStatusUnauthorized(c, "incorrect password"+err.Error())
		return
	}
	response.SendStatusOkWithDataResponse(c, "login succesfully", dto.LoginResponse(user, accessToken))
}
