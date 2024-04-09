package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sky-takeaway/service"
)

func Router() *gin.Engine{
	r := gin.Default()
	r.POST("/admin/login",service.Login)
	r.POST("/admin/register",service.Register)
	return r
}