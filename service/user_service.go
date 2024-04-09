package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sky-takeaway/models"
	"github.com/sky-takeaway/utils"
)

func Login(c *gin.Context) {
	employee := models.Employee{}
	//TODO: 接收json数据
	if err := c.ShouldBindJSON(&employee); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	//TODO: 查询数据库
	sql_employee, db := employee.FindByUserName(employee.UserName)
	if db.Error != nil {
		log.Println(db.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": db.Error.Error(),
			"data":    nil,
		})
		return
	}
	//TODO: md5加密
	password, err := utils.Md5(employee.Password)
	if err!= nil {
		log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    http.StatusBadRequest,
            "message": err.Error(),
            "data":    nil,
        })
        return
	}

	//TODO: 密码比对
	if password != sql_employee.Password{
		log.Println("密码错误")
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    http.StatusBadRequest,
            "message": "密码错误",
            "data":    nil,
        })
        return
	}
	//TODO: 生成token
	token , err := utils.GenerateToken(employee.UserName)
	if err != nil {
		log.Println("token生成失败")
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    http.StatusBadRequest,
            "message": "token生成失败",
            "data":    nil,
        })
        return
	}
	//TODO: 存入redis

	//TODO: 返回token
	c.JSON(http.StatusOK, gin.H{
        "code":    http.StatusOK,
        "message": "登录成功",
        "data":    token,
    })
}

func Register(c *gin.Context) {
	employee := models.Employee{}
	//TODO: 接收json数据
	if err := c.ShouldBindJSON(&employee); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	//TODO: md5加密
	password, err := utils.Md5(employee.Password)
	if err!= nil {
		log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    http.StatusBadRequest,
            "message": err.Error(),
            "data":    nil,
        })
        return
	}
	employee.Password = password
	//TODO: 插入数据库
	employee.Insert(&employee)

	//TODO: 成功
	c.JSON(http.StatusOK, gin.H{
        "code":    http.StatusOK,
        "message": "注册成功",
        "data":    nil,
    })
}
