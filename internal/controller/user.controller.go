package controller

import (
	"github.com/gin-gonic/gin"
	"go-social-network/internal/service"
	"go-social-network/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserInfo(c *gin.Context) {
	response.SuccessResponse(c, 2001, service.NewUserService().GetUserInfo())
}
