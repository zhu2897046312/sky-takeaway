package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sky-takeaway/service"
)

func Router() *gin.Engine{
	r := gin.Default()
	r.POST("/admin/login",service.Login)
	r.POST("/admin/register",service.Register)
	r.GET("/admin/employee/page",service.PageQuery)
	r.POST("/admin/employee/update",service.UpdateEmployee)

	r.GET("/admin/category/page",service.PageQueryCategory)
	r.POST("/admin/category/addCategory",service.AddCategory)
	r.POST("/admin/category/deleteCategory",service.DeleteCategory)
	r.POST("/admin/category/updateCategory",service.UpdateCategory)
	return r
}