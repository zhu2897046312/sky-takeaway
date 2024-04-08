package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sky-takeaway/service"
)

func Router() *gin.Engine{
	r := gin.Default()
	r.POST("/user/login",service.Login)
	r.POST("/user/register",service.Register)
	return r
}