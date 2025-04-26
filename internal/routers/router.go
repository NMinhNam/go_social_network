package routers

import (
	"github.com/gin-gonic/gin"
	"go-social-network/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", controller.NewUserController().GetUserInfo)
	}

	return r
}
